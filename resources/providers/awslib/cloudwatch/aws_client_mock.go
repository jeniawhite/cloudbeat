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

// Code generated by mockery v2.15.0. DO NOT EDIT.

package cloudwatch

import (
	context "context"

	servicecloudwatch "github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	mock "github.com/stretchr/testify/mock"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

// DescribeAlarms provides a mock function with given fields: ctx, params, optFns
func (_m *MockClient) DescribeAlarms(ctx context.Context, params *servicecloudwatch.DescribeAlarmsInput, optFns ...func(*servicecloudwatch.Options)) (*servicecloudwatch.DescribeAlarmsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *servicecloudwatch.DescribeAlarmsOutput
	if rf, ok := ret.Get(0).(func(context.Context, *servicecloudwatch.DescribeAlarmsInput, ...func(*servicecloudwatch.Options)) *servicecloudwatch.DescribeAlarmsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*servicecloudwatch.DescribeAlarmsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *servicecloudwatch.DescribeAlarmsInput, ...func(*servicecloudwatch.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockClient(t mockConstructorTestingTNewMockClient) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
