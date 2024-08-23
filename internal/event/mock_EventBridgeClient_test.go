// Code generated by mockery v2.42.2. DO NOT EDIT.

package event

import (
	context "context"

	eventbridge "github.com/aws/aws-sdk-go-v2/service/eventbridge"
	mock "github.com/stretchr/testify/mock"
)

// mockEventBridgeClient is an autogenerated mock type for the EventBridgeClient type
type mockEventBridgeClient struct {
	mock.Mock
}

type mockEventBridgeClient_Expecter struct {
	mock *mock.Mock
}

func (_m *mockEventBridgeClient) EXPECT() *mockEventBridgeClient_Expecter {
	return &mockEventBridgeClient_Expecter{mock: &_m.Mock}
}

// PutEvents provides a mock function with given fields: ctx, params, optFns
func (_m *mockEventBridgeClient) PutEvents(ctx context.Context, params *eventbridge.PutEventsInput, optFns ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutEvents")
	}

	var r0 *eventbridge.PutEventsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *eventbridge.PutEventsInput, ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *eventbridge.PutEventsInput, ...func(*eventbridge.Options)) *eventbridge.PutEventsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*eventbridge.PutEventsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *eventbridge.PutEventsInput, ...func(*eventbridge.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mockEventBridgeClient_PutEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutEvents'
type mockEventBridgeClient_PutEvents_Call struct {
	*mock.Call
}

// PutEvents is a helper method to define mock.On call
//   - ctx context.Context
//   - params *eventbridge.PutEventsInput
//   - optFns ...func(*eventbridge.Options)
func (_e *mockEventBridgeClient_Expecter) PutEvents(ctx interface{}, params interface{}, optFns ...interface{}) *mockEventBridgeClient_PutEvents_Call {
	return &mockEventBridgeClient_PutEvents_Call{Call: _e.mock.On("PutEvents",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *mockEventBridgeClient_PutEvents_Call) Run(run func(ctx context.Context, params *eventbridge.PutEventsInput, optFns ...func(*eventbridge.Options))) *mockEventBridgeClient_PutEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*eventbridge.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*eventbridge.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*eventbridge.PutEventsInput), variadicArgs...)
	})
	return _c
}

func (_c *mockEventBridgeClient_PutEvents_Call) Return(_a0 *eventbridge.PutEventsOutput, _a1 error) *mockEventBridgeClient_PutEvents_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *mockEventBridgeClient_PutEvents_Call) RunAndReturn(run func(context.Context, *eventbridge.PutEventsInput, ...func(*eventbridge.Options)) (*eventbridge.PutEventsOutput, error)) *mockEventBridgeClient_PutEvents_Call {
	_c.Call.Return(run)
	return _c
}

// newMockEventBridgeClient creates a new instance of mockEventBridgeClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockEventBridgeClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockEventBridgeClient {
	mock := &mockEventBridgeClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
