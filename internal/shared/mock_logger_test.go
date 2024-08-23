// Code generated by mockery v2.42.2. DO NOT EDIT.

package shared

import mock "github.com/stretchr/testify/mock"

// mockLogger is an autogenerated mock type for the logger type
type mockLogger struct {
	mock.Mock
}

type mockLogger_Expecter struct {
	mock *mock.Mock
}

func (_m *mockLogger) EXPECT() *mockLogger_Expecter {
	return &mockLogger_Expecter{mock: &_m.Mock}
}

// Error provides a mock function with given fields: _a0, _a1
func (_m *mockLogger) Error(_a0 string, _a1 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	_m.Called(_ca...)
}

// mockLogger_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type mockLogger_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
//   - _a0 string
//   - _a1 ...interface{}
func (_e *mockLogger_Expecter) Error(_a0 interface{}, _a1 ...interface{}) *mockLogger_Error_Call {
	return &mockLogger_Error_Call{Call: _e.mock.On("Error",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *mockLogger_Error_Call) Run(run func(_a0 string, _a1 ...interface{})) *mockLogger_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *mockLogger_Error_Call) Return() *mockLogger_Error_Call {
	_c.Call.Return()
	return _c
}

func (_c *mockLogger_Error_Call) RunAndReturn(run func(string, ...interface{})) *mockLogger_Error_Call {
	_c.Call.Return(run)
	return _c
}

// Info provides a mock function with given fields: _a0, _a1
func (_m *mockLogger) Info(_a0 string, _a1 ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	_m.Called(_ca...)
}

// mockLogger_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type mockLogger_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
//   - _a0 string
//   - _a1 ...interface{}
func (_e *mockLogger_Expecter) Info(_a0 interface{}, _a1 ...interface{}) *mockLogger_Info_Call {
	return &mockLogger_Info_Call{Call: _e.mock.On("Info",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *mockLogger_Info_Call) Run(run func(_a0 string, _a1 ...interface{})) *mockLogger_Info_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *mockLogger_Info_Call) Return() *mockLogger_Info_Call {
	_c.Call.Return()
	return _c
}

func (_c *mockLogger_Info_Call) RunAndReturn(run func(string, ...interface{})) *mockLogger_Info_Call {
	_c.Call.Return(run)
	return _c
}

// newMockLogger creates a new instance of mockLogger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockLogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockLogger {
	mock := &mockLogger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
