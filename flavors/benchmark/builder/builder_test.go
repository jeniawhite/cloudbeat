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

package builder

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/elastic/cloudbeat/config"
	"github.com/elastic/cloudbeat/dataprovider"
	"github.com/elastic/cloudbeat/resources/fetching"
	"github.com/elastic/cloudbeat/resources/fetching/registry"
	"github.com/elastic/cloudbeat/resources/utils/testhelper"
	"github.com/elastic/cloudbeat/uniqueness"
)

func TestBase_Build_Success(t *testing.T) {
	tests := []struct {
		name      string
		opts      []Option
		benchType interface{}
	}{
		{
			name:      "by default create base benchmark",
			benchType: &basebenchmark{}, //nolint:exhaustruct
		},
		{
			name: "with opts create base benchmark",
			opts: []Option{
				WithIdProvider(dataprovider.NewMockIdProvider(t)),
				WithManagerTimeout(time.Minute),
				WithBenchmarkDataProvider(dataprovider.NewMockCommonDataProvider(t)),
			},
			benchType: &basebenchmark{}, //nolint:exhaustruct
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := testhelper.NewLogger(t)
			path, err := filepath.Abs("../../../bundle.tar.gz")
			assert.NoError(t, err)

			resourceCh := make(chan fetching.ResourceInfo)
			reg := registry.NewMockRegistry(t)
			benchmark, err := New(tt.opts...).Build(context.Background(), log, &config.Config{
				BundlePath: path,
				Period:     time.Minute,
			}, resourceCh, reg)
			assert.NoError(t, err)
			assert.IsType(t, tt.benchType, benchmark)

			reg.EXPECT().Keys().Return([]string{}).Twice()
			reg.EXPECT().Update().Return().Once()
			_, err = benchmark.Run(context.Background())
			time.Sleep(100 * time.Millisecond)
			assert.NoError(t, err)
		})
	}
}

func TestBase_BuildK8s_Success(t *testing.T) {
	tests := []struct {
		name      string
		opts      []Option
		benchType interface{}
	}{
		{
			name:      "by default create k8s benchmark",
			benchType: &k8sbenchmark{}, //nolint:exhaustruct

		}, {
			name: "with opts create k8s benchmark",
			opts: []Option{
				WithIdProvider(dataprovider.NewMockIdProvider(t)),
				WithManagerTimeout(time.Minute),
				WithBenchmarkDataProvider(dataprovider.NewMockCommonDataProvider(t)),
			},
			benchType: &k8sbenchmark{}, //nolint:exhaustruct
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log := testhelper.NewLogger(t)
			path, err := filepath.Abs("../../../bundle.tar.gz")
			assert.NoError(t, err)

			resourceCh := make(chan fetching.ResourceInfo)
			reg := registry.NewMockRegistry(t)
			le := uniqueness.NewMockManager(t)
			benchmark, err := New(tt.opts...).BuildK8s(context.Background(), log, &config.Config{
				BundlePath: path,
				Period:     time.Minute,
			}, resourceCh, reg, le)
			assert.NoError(t, err)
			assert.IsType(t, tt.benchType, benchmark)

			reg.EXPECT().Keys().Return([]string{}).Twice()
			reg.EXPECT().Update().Return().Once()
			le.EXPECT().Run(mock.Anything).Return(nil).Once()
			_, err = benchmark.Run(context.Background())
			time.Sleep(100 * time.Millisecond)
			assert.NoError(t, err)
		})
	}
}