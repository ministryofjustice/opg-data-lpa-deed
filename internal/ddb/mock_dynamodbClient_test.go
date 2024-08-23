// Code generated by mockery v2.42.2. DO NOT EDIT.

package ddb

import (
	context "context"

	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	mock "github.com/stretchr/testify/mock"
)

// mockDynamodbClient is an autogenerated mock type for the dynamodbClient type
type mockDynamodbClient struct {
	mock.Mock
}

type mockDynamodbClient_Expecter struct {
	mock *mock.Mock
}

func (_m *mockDynamodbClient) EXPECT() *mockDynamodbClient_Expecter {
	return &mockDynamodbClient_Expecter{mock: &_m.Mock}
}

// BatchGetItem provides a mock function with given fields: ctx, params, optFns
func (_m *mockDynamodbClient) BatchGetItem(ctx context.Context, params *dynamodb.BatchGetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BatchGetItem")
	}

	var r0 *dynamodb.BatchGetItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.BatchGetItemInput, ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.BatchGetItemInput, ...func(*dynamodb.Options)) *dynamodb.BatchGetItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.BatchGetItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.BatchGetItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockDynamodbClient_BatchGetItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchGetItem'
type mockDynamodbClient_BatchGetItem_Call struct {
	*mock.Call
}

// BatchGetItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.BatchGetItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *mockDynamodbClient_Expecter) BatchGetItem(ctx interface{}, params interface{}, optFns ...interface{}) *mockDynamodbClient_BatchGetItem_Call {
	return &mockDynamodbClient_BatchGetItem_Call{Call: _e.mock.On("BatchGetItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *mockDynamodbClient_BatchGetItem_Call) Run(run func(ctx context.Context, params *dynamodb.BatchGetItemInput, optFns ...func(*dynamodb.Options))) *mockDynamodbClient_BatchGetItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.BatchGetItemInput), variadicArgs...)
	})
	return _c
}

func (_c *mockDynamodbClient_BatchGetItem_Call) Return(_a0 *dynamodb.BatchGetItemOutput, _a1 error) *mockDynamodbClient_BatchGetItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockDynamodbClient_BatchGetItem_Call) RunAndReturn(run func(context.Context, *dynamodb.BatchGetItemInput, ...func(*dynamodb.Options)) (*dynamodb.BatchGetItemOutput, error)) *mockDynamodbClient_BatchGetItem_Call {
	_c.Call.Return(run)
	return _c
}

// GetItem provides a mock function with given fields: ctx, params, optFns
func (_m *mockDynamodbClient) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetItem")
	}

	var r0 *dynamodb.GetItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) *dynamodb.GetItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.GetItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockDynamodbClient_GetItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetItem'
type mockDynamodbClient_GetItem_Call struct {
	*mock.Call
}

// GetItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.GetItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *mockDynamodbClient_Expecter) GetItem(ctx interface{}, params interface{}, optFns ...interface{}) *mockDynamodbClient_GetItem_Call {
	return &mockDynamodbClient_GetItem_Call{Call: _e.mock.On("GetItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *mockDynamodbClient_GetItem_Call) Run(run func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options))) *mockDynamodbClient_GetItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.GetItemInput), variadicArgs...)
	})
	return _c
}

func (_c *mockDynamodbClient_GetItem_Call) Return(_a0 *dynamodb.GetItemOutput, _a1 error) *mockDynamodbClient_GetItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockDynamodbClient_GetItem_Call) RunAndReturn(run func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)) *mockDynamodbClient_GetItem_Call {
	_c.Call.Return(run)
	return _c
}

// PutItem provides a mock function with given fields: ctx, params, optFns
func (_m *mockDynamodbClient) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutItem")
	}

	var r0 *dynamodb.PutItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) *dynamodb.PutItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.PutItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockDynamodbClient_PutItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutItem'
type mockDynamodbClient_PutItem_Call struct {
	*mock.Call
}

// PutItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.PutItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *mockDynamodbClient_Expecter) PutItem(ctx interface{}, params interface{}, optFns ...interface{}) *mockDynamodbClient_PutItem_Call {
	return &mockDynamodbClient_PutItem_Call{Call: _e.mock.On("PutItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *mockDynamodbClient_PutItem_Call) Run(run func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options))) *mockDynamodbClient_PutItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.PutItemInput), variadicArgs...)
	})
	return _c
}

func (_c *mockDynamodbClient_PutItem_Call) Return(_a0 *dynamodb.PutItemOutput, _a1 error) *mockDynamodbClient_PutItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockDynamodbClient_PutItem_Call) RunAndReturn(run func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)) *mockDynamodbClient_PutItem_Call {
	_c.Call.Return(run)
	return _c
}

// TransactWriteItems provides a mock function with given fields: ctx, params, optFns
func (_m *mockDynamodbClient) TransactWriteItems(ctx context.Context, params *dynamodb.TransactWriteItemsInput, optFns ...func(*dynamodb.Options)) (*dynamodb.TransactWriteItemsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TransactWriteItems")
	}

	var r0 *dynamodb.TransactWriteItemsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.TransactWriteItemsInput, ...func(*dynamodb.Options)) (*dynamodb.TransactWriteItemsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.TransactWriteItemsInput, ...func(*dynamodb.Options)) *dynamodb.TransactWriteItemsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.TransactWriteItemsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.TransactWriteItemsInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockDynamodbClient_TransactWriteItems_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TransactWriteItems'
type mockDynamodbClient_TransactWriteItems_Call struct {
	*mock.Call
}

// TransactWriteItems is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.TransactWriteItemsInput
//   - optFns ...func(*dynamodb.Options)
func (_e *mockDynamodbClient_Expecter) TransactWriteItems(ctx interface{}, params interface{}, optFns ...interface{}) *mockDynamodbClient_TransactWriteItems_Call {
	return &mockDynamodbClient_TransactWriteItems_Call{Call: _e.mock.On("TransactWriteItems",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *mockDynamodbClient_TransactWriteItems_Call) Run(run func(ctx context.Context, params *dynamodb.TransactWriteItemsInput, optFns ...func(*dynamodb.Options))) *mockDynamodbClient_TransactWriteItems_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.TransactWriteItemsInput), variadicArgs...)
	})
	return _c
}

func (_c *mockDynamodbClient_TransactWriteItems_Call) Return(_a0 *dynamodb.TransactWriteItemsOutput, _a1 error) *mockDynamodbClient_TransactWriteItems_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockDynamodbClient_TransactWriteItems_Call) RunAndReturn(run func(context.Context, *dynamodb.TransactWriteItemsInput, ...func(*dynamodb.Options)) (*dynamodb.TransactWriteItemsOutput, error)) *mockDynamodbClient_TransactWriteItems_Call {
	_c.Call.Return(run)
	return _c
}

// newMockDynamodbClient creates a new instance of mockDynamodbClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockDynamodbClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockDynamodbClient {
	mock := &mockDynamodbClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
