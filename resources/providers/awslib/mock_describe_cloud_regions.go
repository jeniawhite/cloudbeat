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

// Code generated by mockery v2.33.3. DO NOT EDIT.

package awslib

import (
	context "context"

	ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	mock "github.com/stretchr/testify/mock"
)

// mockDescribeCloudRegions is an autogenerated mock type for the describeCloudRegions type
type mockDescribeCloudRegions struct {
	mock.Mock
}

type mockDescribeCloudRegions_Expecter struct {
	mock *mock.Mock
}

func (_m *mockDescribeCloudRegions) EXPECT() *mockDescribeCloudRegions_Expecter {
	return &mockDescribeCloudRegions_Expecter{mock: &_m.Mock}
}

// DescribeRegions provides a mock function with given fields: ctx, params, optFns
func (_m *mockDescribeCloudRegions) DescribeRegions(ctx context.Context, params *ec2.DescribeRegionsInput, optFns ...func(*ec2.Options)) (*ec2.DescribeRegionsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ec2.DescribeRegionsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *ec2.DescribeRegionsInput, ...func(*ec2.Options)) (*ec2.DescribeRegionsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *ec2.DescribeRegionsInput, ...func(*ec2.Options)) *ec2.DescribeRegionsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ec2.DescribeRegionsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *ec2.DescribeRegionsInput, ...func(*ec2.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockDescribeCloudRegions_DescribeRegions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeRegions'
type mockDescribeCloudRegions_DescribeRegions_Call struct {
	*mock.Call
}

// DescribeRegions is a helper method to define mock.On call
//   - ctx context.Context
//   - params *ec2.DescribeRegionsInput
//   - optFns ...func(*ec2.Options)
func (_e *mockDescribeCloudRegions_Expecter) DescribeRegions(ctx interface{}, params interface{}, optFns ...interface{}) *mockDescribeCloudRegions_DescribeRegions_Call {
	return &mockDescribeCloudRegions_DescribeRegions_Call{Call: _e.mock.On("DescribeRegions",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *mockDescribeCloudRegions_DescribeRegions_Call) Run(run func(ctx context.Context, params *ec2.DescribeRegionsInput, optFns ...func(*ec2.Options))) *mockDescribeCloudRegions_DescribeRegions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*ec2.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*ec2.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*ec2.DescribeRegionsInput), variadicArgs...)
	})
	return _c
}

func (_c *mockDescribeCloudRegions_DescribeRegions_Call) Return(_a0 *ec2.DescribeRegionsOutput, _a1 error) *mockDescribeCloudRegions_DescribeRegions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockDescribeCloudRegions_DescribeRegions_Call) RunAndReturn(run func(context.Context, *ec2.DescribeRegionsInput, ...func(*ec2.Options)) (*ec2.DescribeRegionsOutput, error)) *mockDescribeCloudRegions_DescribeRegions_Call {
	_c.Call.Return(run)
	return _c
}

// newMockDescribeCloudRegions creates a new instance of mockDescribeCloudRegions. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockDescribeCloudRegions(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockDescribeCloudRegions {
	mock := &mockDescribeCloudRegions{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
