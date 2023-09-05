// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package benchmark

import (
	"context"
	"fmt"
	"testing"

	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/elastic/beats/v7/x-pack/libbeat/common/aws"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
	core_v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sfake "k8s.io/client-go/kubernetes/fake"

	"github.com/elastic/cloudbeat/config"
	"github.com/elastic/cloudbeat/dataprovider/providers/cloud"
	"github.com/elastic/cloudbeat/dataprovider/providers/k8s"
	"github.com/elastic/cloudbeat/resources/fetching"
	"github.com/elastic/cloudbeat/resources/providers/awslib"
	"github.com/elastic/cloudbeat/resources/utils/testhelper"
)

func TestNewBenchmark(t *testing.T) {
	defer goleak.VerifyNone(t, goleak.IgnoreCurrent()) // NewBenchmark should not start anything

	tests := []struct {
		cfg      config.Config
		wantType Benchmark
		wantErr  bool
	}{
		{
			cfg:     config.Config{Benchmark: "unknown"},
			wantErr: true,
		},
		{
			cfg:      config.Config{Benchmark: config.CIS_AWS},
			wantType: &AWS{}, //nolint:exhaustruct
		},
		{
			cfg: config.Config{
				Benchmark: config.CIS_AWS,
				CloudConfig: config.CloudConfig{
					Aws: config.AwsConfig{AccountType: config.SingleAccount},
				},
			},
			wantType: &AWS{}, //nolint:exhaustruct
		},
		{
			cfg: config.Config{
				Benchmark: config.CIS_AWS,
				CloudConfig: config.CloudConfig{
					Aws: config.AwsConfig{AccountType: config.OrganizationAccount},
				},
			},
			wantType: &AWSOrg{}, //nolint:exhaustruct
		},
		{
			cfg:      config.Config{Benchmark: config.CIS_EKS},
			wantType: &EKS{}, //nolint:exhaustruct
		},
		{
			cfg:      config.Config{Benchmark: config.CIS_K8S},
			wantType: &K8S{}, //nolint:exhaustruct
		},
		{
			cfg:      config.Config{Benchmark: config.CIS_GCP},
			wantType: &GCP{}, //nolint:exhaustruct
		},
		// TODO: Uncomment when azure.go and configuration is merged
		// {
		// 	cfg:      config.Config{Benchmark: config.CIS_AZURE},
		// 	wantType: &Azure{}, //nolint:exhaustruct
		// },
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%T", tt.wantType), func(t *testing.T) {
			got, err := NewBenchmark(&tt.cfg)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.IsType(t, tt.wantType, got)
			assert.NoError(t, got.checkDependencies())
		})
	}
}

func testInitialize(t *testing.T, benchmark Benchmark, cfg *config.Config, wantErr string, want []string) {
	t.Helper()

	reg, dp, err := benchmark.Initialize(context.Background(), testhelper.NewLogger(t), cfg, make(chan fetching.ResourceInfo))
	if wantErr != "" {
		assert.ErrorContains(t, err, wantErr)
		return
	}

	require.NoError(t, err)
	assert.Len(t, reg.Keys(), len(want))

	require.NoError(t, benchmark.Run(context.Background()))
	defer benchmark.Stop()

	for _, fetcher := range want {
		ok := reg.ShouldRun(fetcher)
		assert.Truef(t, ok, "fetcher %s enabled", fetcher)
	}

	// TODO: gcp diff tests cover
	assert.NotNil(t, dp)
}

func mockAwsCfg(err error) *awslib.MockConfigProviderAPI {
	awsCfg := awslib.MockConfigProviderAPI{}
	awsCfg.EXPECT().InitializeAWSConfig(mock.Anything, mock.Anything).
		Call.
		Return(
			func(ctx context.Context, config aws.ConfigAWS) *awssdk.Config {
				if err != nil {
					return nil
				}

				awsConfig := awssdk.NewConfig()
				awsCredentials := awssdk.Credentials{
					AccessKeyID:     config.AccessKeyID,
					SecretAccessKey: config.SecretAccessKey,
					SessionToken:    config.SessionToken,
				}

				awsConfig.Credentials = credentials.StaticCredentialsProvider{
					Value: awsCredentials,
				}
				awsConfig.Region = "us1-east"
				return awsConfig
			},
			func(ctx context.Context, config aws.ConfigAWS) error {
				return err
			},
		)
	return &awsCfg
}

func mockKubeClient(err error) k8s.ClientGetterAPI {
	kube := k8s.MockClientGetterAPI{}
	on := kube.EXPECT().GetClient(mock.Anything, mock.Anything, mock.Anything)
	if err == nil {
		on.Return(
			k8sfake.NewSimpleClientset(
				&core_v1.Node{
					ObjectMeta: metav1.ObjectMeta{
						Name: "node-name",
					},
				},
				&core_v1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: "kube-system",
					},
				},
			), nil)
	} else {
		on.Return(nil, err)
	}
	return &kube
}

func mockAwsIdentityProvider(err error) *awslib.MockIdentityProviderGetter {
	identityProvider := &awslib.MockIdentityProviderGetter{}
	on := identityProvider.EXPECT().GetIdentity(mock.Anything, mock.Anything)
	if err == nil {
		on.Return(
			&cloud.Identity{
				Account: "test-account",
			},
			nil,
		)
	} else {
		on.Return(nil, err)
	}
	return identityProvider
}
