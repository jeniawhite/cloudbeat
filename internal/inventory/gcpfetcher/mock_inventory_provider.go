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

// Code generated by mockery v2.37.1. DO NOT EDIT.

package gcpfetcher

import (
	context "context"

	inventory "github.com/elastic/cloudbeat/internal/resources/providers/gcplib/inventory"
	mock "github.com/stretchr/testify/mock"
)

// mockInventoryProvider is an autogenerated mock type for the inventoryProvider type
type mockInventoryProvider struct {
	mock.Mock
}

type mockInventoryProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *mockInventoryProvider) EXPECT() *mockInventoryProvider_Expecter {
	return &mockInventoryProvider_Expecter{mock: &_m.Mock}
}

// ListAllAssetTypesByName provides a mock function with given fields: ctx, assets
func (_m *mockInventoryProvider) ListAllAssetTypesByName(ctx context.Context, assets []string) ([]*inventory.ExtendedGcpAsset, error) {
	ret := _m.Called(ctx, assets)

	var r0 []*inventory.ExtendedGcpAsset
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]*inventory.ExtendedGcpAsset, error)); ok {
		return rf(ctx, assets)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*inventory.ExtendedGcpAsset); ok {
		r0 = rf(ctx, assets)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*inventory.ExtendedGcpAsset)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, assets)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockInventoryProvider_ListAllAssetTypesByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAllAssetTypesByName'
type mockInventoryProvider_ListAllAssetTypesByName_Call struct {
	*mock.Call
}

// ListAllAssetTypesByName is a helper method to define mock.On call
//   - ctx context.Context
//   - assets []string
func (_e *mockInventoryProvider_Expecter) ListAllAssetTypesByName(ctx interface{}, assets interface{}) *mockInventoryProvider_ListAllAssetTypesByName_Call {
	return &mockInventoryProvider_ListAllAssetTypesByName_Call{Call: _e.mock.On("ListAllAssetTypesByName", ctx, assets)}
}

func (_c *mockInventoryProvider_ListAllAssetTypesByName_Call) Run(run func(ctx context.Context, assets []string)) *mockInventoryProvider_ListAllAssetTypesByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]string))
	})
	return _c
}

func (_c *mockInventoryProvider_ListAllAssetTypesByName_Call) Return(_a0 []*inventory.ExtendedGcpAsset, _a1 error) *mockInventoryProvider_ListAllAssetTypesByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockInventoryProvider_ListAllAssetTypesByName_Call) RunAndReturn(run func(context.Context, []string) ([]*inventory.ExtendedGcpAsset, error)) *mockInventoryProvider_ListAllAssetTypesByName_Call {
	_c.Call.Return(run)
	return _c
}

// newMockInventoryProvider creates a new instance of mockInventoryProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockInventoryProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockInventoryProvider {
	mock := &mockInventoryProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}