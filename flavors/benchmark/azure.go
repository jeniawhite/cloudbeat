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
	"errors"
	"fmt"

	"github.com/elastic/elastic-agent-libs/logp"

	"github.com/elastic/cloudbeat/config"
	"github.com/elastic/cloudbeat/dataprovider"
	"github.com/elastic/cloudbeat/dataprovider/providers/cloud"
	"github.com/elastic/cloudbeat/resources/fetching"
	"github.com/elastic/cloudbeat/resources/fetching/factory"
	"github.com/elastic/cloudbeat/resources/fetching/registry"
	"github.com/elastic/cloudbeat/resources/providers/azurelib/auth"
	"github.com/elastic/cloudbeat/resources/providers/azurelib/inventory"
)

type Azure struct {
	// IdentityProvider     identity.ProviderGetter
	CfgProvider          auth.ConfigProviderAPI
	inventoryInitializer inventory.ProviderInitializerAPI
}

func (a *Azure) Run(context.Context) error { return nil }

func (a *Azure) Initialize(ctx context.Context, log *logp.Logger, cfg *config.Config, ch chan fetching.ResourceInfo) (registry.Registry, dataprovider.CommonDataProvider, dataprovider.IdProvider, error) {
	if err := a.checkDependencies(); err != nil {
		return nil, nil, nil, err
	}

	azureConfig, err := a.CfgProvider.GetAzureClientConfig(cfg.CloudConfig.Azure, log)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to initialize azure config: %w", err)
	}

	// azureIdentity, err := a.IdentityProvider.GetIdentity(ctx, azureConfig)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("failed to get Azure identity: %v", err)
	// }

	assetProvider, err := a.inventoryInitializer.Init(ctx, log, *azureConfig)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to initialize azure asset inventory: %v", err)
	}

	fetchers, err := factory.NewCisAzureFactory(ctx, log, ch, assetProvider)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to initialize azure fetchers: %v", err)
	}

	return registry.NewRegistry(log, fetchers), cloud.NewDataProvider(
		cloud.WithLogger(log),
		// cloud.WithAccount(*azureIdentity),
	), cloud.NewIdProvider(), nil
}

func (a *Azure) Stop() {}

func (a *Azure) checkDependencies() error {
	// if a.IdentityProvider == nil {
	// 	return errors.New("azure identity provider is uninitialized")
	// }

	if a.CfgProvider == nil {
		return errors.New("azure config provider is uninitialized")
	}

	if a.inventoryInitializer == nil {
		return errors.New("azure asset inventory is uninitialized")
	}
	return nil
}