// Code generated by mockery v2.23.4. DO NOT EDIT.

package payment

import (
	model "github.com/kevindharmawan/saas-backend/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// MockPaymentRepository is an autogenerated mock type for the PaymentRepository type
type MockPaymentRepository struct {
	mock.Mock
}

type MockPaymentRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPaymentRepository) EXPECT() *MockPaymentRepository_Expecter {
	return &MockPaymentRepository_Expecter{mock: &_m.Mock}
}

// AddCreditCard provides a mock function with given fields: customerId, cardToken
func (_m *MockPaymentRepository) AddCreditCard(customerId string, cardToken string) (*model.Card, *model.AppError) {
	ret := _m.Called(customerId, cardToken)

	var r0 *model.Card
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string, string) (*model.Card, *model.AppError)); ok {
		return rf(customerId, cardToken)
	}
	if rf, ok := ret.Get(0).(func(string, string) *model.Card); ok {
		r0 = rf(customerId, cardToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Card)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) *model.AppError); ok {
		r1 = rf(customerId, cardToken)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// MockPaymentRepository_AddCreditCard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddCreditCard'
type MockPaymentRepository_AddCreditCard_Call struct {
	*mock.Call
}

// AddCreditCard is a helper method to define mock.On call
//   - customerId string
//   - cardToken string
func (_e *MockPaymentRepository_Expecter) AddCreditCard(customerId interface{}, cardToken interface{}) *MockPaymentRepository_AddCreditCard_Call {
	return &MockPaymentRepository_AddCreditCard_Call{Call: _e.mock.On("AddCreditCard", customerId, cardToken)}
}

func (_c *MockPaymentRepository_AddCreditCard_Call) Run(run func(customerId string, cardToken string)) *MockPaymentRepository_AddCreditCard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockPaymentRepository_AddCreditCard_Call) Return(_a0 *model.Card, _a1 *model.AppError) *MockPaymentRepository_AddCreditCard_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPaymentRepository_AddCreditCard_Call) RunAndReturn(run func(string, string) (*model.Card, *model.AppError)) *MockPaymentRepository_AddCreditCard_Call {
	_c.Call.Return(run)
	return _c
}

// CancelSubscription provides a mock function with given fields: customerId
func (_m *MockPaymentRepository) CancelSubscription(customerId string) *model.AppError {
	ret := _m.Called(customerId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(customerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// MockPaymentRepository_CancelSubscription_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelSubscription'
type MockPaymentRepository_CancelSubscription_Call struct {
	*mock.Call
}

// CancelSubscription is a helper method to define mock.On call
//   - customerId string
func (_e *MockPaymentRepository_Expecter) CancelSubscription(customerId interface{}) *MockPaymentRepository_CancelSubscription_Call {
	return &MockPaymentRepository_CancelSubscription_Call{Call: _e.mock.On("CancelSubscription", customerId)}
}

func (_c *MockPaymentRepository_CancelSubscription_Call) Run(run func(customerId string)) *MockPaymentRepository_CancelSubscription_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPaymentRepository_CancelSubscription_Call) Return(_a0 *model.AppError) *MockPaymentRepository_CancelSubscription_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPaymentRepository_CancelSubscription_Call) RunAndReturn(run func(string) *model.AppError) *MockPaymentRepository_CancelSubscription_Call {
	_c.Call.Return(run)
	return _c
}

// CreateCustomer provides a mock function with given fields: user
func (_m *MockPaymentRepository) CreateCustomer(user *model.User) (*model.User, *model.AppError) {
	ret := _m.Called(user)

	var r0 *model.User
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(*model.User) (*model.User, *model.AppError)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(*model.User) *model.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.User) *model.AppError); ok {
		r1 = rf(user)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// MockPaymentRepository_CreateCustomer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCustomer'
type MockPaymentRepository_CreateCustomer_Call struct {
	*mock.Call
}

// CreateCustomer is a helper method to define mock.On call
//   - user *model.User
func (_e *MockPaymentRepository_Expecter) CreateCustomer(user interface{}) *MockPaymentRepository_CreateCustomer_Call {
	return &MockPaymentRepository_CreateCustomer_Call{Call: _e.mock.On("CreateCustomer", user)}
}

func (_c *MockPaymentRepository_CreateCustomer_Call) Run(run func(user *model.User)) *MockPaymentRepository_CreateCustomer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.User))
	})
	return _c
}

func (_c *MockPaymentRepository_CreateCustomer_Call) Return(_a0 *model.User, _a1 *model.AppError) *MockPaymentRepository_CreateCustomer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPaymentRepository_CreateCustomer_Call) RunAndReturn(run func(*model.User) (*model.User, *model.AppError)) *MockPaymentRepository_CreateCustomer_Call {
	_c.Call.Return(run)
	return _c
}

// GetCreditCards provides a mock function with given fields: customerId
func (_m *MockPaymentRepository) GetCreditCards(customerId string) ([]*model.Card, *model.AppError) {
	ret := _m.Called(customerId)

	var r0 []*model.Card
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string) ([]*model.Card, *model.AppError)); ok {
		return rf(customerId)
	}
	if rf, ok := ret.Get(0).(func(string) []*model.Card); ok {
		r0 = rf(customerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Card)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(customerId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// MockPaymentRepository_GetCreditCards_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCreditCards'
type MockPaymentRepository_GetCreditCards_Call struct {
	*mock.Call
}

// GetCreditCards is a helper method to define mock.On call
//   - customerId string
func (_e *MockPaymentRepository_Expecter) GetCreditCards(customerId interface{}) *MockPaymentRepository_GetCreditCards_Call {
	return &MockPaymentRepository_GetCreditCards_Call{Call: _e.mock.On("GetCreditCards", customerId)}
}

func (_c *MockPaymentRepository_GetCreditCards_Call) Run(run func(customerId string)) *MockPaymentRepository_GetCreditCards_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPaymentRepository_GetCreditCards_Call) Return(_a0 []*model.Card, _a1 *model.AppError) *MockPaymentRepository_GetCreditCards_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPaymentRepository_GetCreditCards_Call) RunAndReturn(run func(string) ([]*model.Card, *model.AppError)) *MockPaymentRepository_GetCreditCards_Call {
	_c.Call.Return(run)
	return _c
}

// GetCustomerByUserId provides a mock function with given fields: userId
func (_m *MockPaymentRepository) GetCustomerByUserId(userId int64) (*model.User, *model.AppError) {
	ret := _m.Called(userId)

	var r0 *model.User
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(int64) (*model.User, *model.AppError)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(int64) *model.User); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) *model.AppError); ok {
		r1 = rf(userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// MockPaymentRepository_GetCustomerByUserId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCustomerByUserId'
type MockPaymentRepository_GetCustomerByUserId_Call struct {
	*mock.Call
}

// GetCustomerByUserId is a helper method to define mock.On call
//   - userId int64
func (_e *MockPaymentRepository_Expecter) GetCustomerByUserId(userId interface{}) *MockPaymentRepository_GetCustomerByUserId_Call {
	return &MockPaymentRepository_GetCustomerByUserId_Call{Call: _e.mock.On("GetCustomerByUserId", userId)}
}

func (_c *MockPaymentRepository_GetCustomerByUserId_Call) Run(run func(userId int64)) *MockPaymentRepository_GetCustomerByUserId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockPaymentRepository_GetCustomerByUserId_Call) Return(_a0 *model.User, _a1 *model.AppError) *MockPaymentRepository_GetCustomerByUserId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPaymentRepository_GetCustomerByUserId_Call) RunAndReturn(run func(int64) (*model.User, *model.AppError)) *MockPaymentRepository_GetCustomerByUserId_Call {
	_c.Call.Return(run)
	return _c
}

// GetSubscriptionInfo provides a mock function with given fields: customerId
func (_m *MockPaymentRepository) GetSubscriptionInfo(customerId string) (*model.Subscription, *model.AppError) {
	ret := _m.Called(customerId)

	var r0 *model.Subscription
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string) (*model.Subscription, *model.AppError)); ok {
		return rf(customerId)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Subscription); ok {
		r0 = rf(customerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(customerId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// MockPaymentRepository_GetSubscriptionInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubscriptionInfo'
type MockPaymentRepository_GetSubscriptionInfo_Call struct {
	*mock.Call
}

// GetSubscriptionInfo is a helper method to define mock.On call
//   - customerId string
func (_e *MockPaymentRepository_Expecter) GetSubscriptionInfo(customerId interface{}) *MockPaymentRepository_GetSubscriptionInfo_Call {
	return &MockPaymentRepository_GetSubscriptionInfo_Call{Call: _e.mock.On("GetSubscriptionInfo", customerId)}
}

func (_c *MockPaymentRepository_GetSubscriptionInfo_Call) Run(run func(customerId string)) *MockPaymentRepository_GetSubscriptionInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPaymentRepository_GetSubscriptionInfo_Call) Return(_a0 *model.Subscription, _a1 *model.AppError) *MockPaymentRepository_GetSubscriptionInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPaymentRepository_GetSubscriptionInfo_Call) RunAndReturn(run func(string) (*model.Subscription, *model.AppError)) *MockPaymentRepository_GetSubscriptionInfo_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveCreditCard provides a mock function with given fields: customerId, cardId
func (_m *MockPaymentRepository) RemoveCreditCard(customerId string, cardId string) *model.AppError {
	ret := _m.Called(customerId, cardId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, string) *model.AppError); ok {
		r0 = rf(customerId, cardId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// MockPaymentRepository_RemoveCreditCard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveCreditCard'
type MockPaymentRepository_RemoveCreditCard_Call struct {
	*mock.Call
}

// RemoveCreditCard is a helper method to define mock.On call
//   - customerId string
//   - cardId string
func (_e *MockPaymentRepository_Expecter) RemoveCreditCard(customerId interface{}, cardId interface{}) *MockPaymentRepository_RemoveCreditCard_Call {
	return &MockPaymentRepository_RemoveCreditCard_Call{Call: _e.mock.On("RemoveCreditCard", customerId, cardId)}
}

func (_c *MockPaymentRepository_RemoveCreditCard_Call) Run(run func(customerId string, cardId string)) *MockPaymentRepository_RemoveCreditCard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockPaymentRepository_RemoveCreditCard_Call) Return(_a0 *model.AppError) *MockPaymentRepository_RemoveCreditCard_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPaymentRepository_RemoveCreditCard_Call) RunAndReturn(run func(string, string) *model.AppError) *MockPaymentRepository_RemoveCreditCard_Call {
	_c.Call.Return(run)
	return _c
}

// SetDefaultCreditCard provides a mock function with given fields: customerId, cardId
func (_m *MockPaymentRepository) SetDefaultCreditCard(customerId string, cardId string) *model.AppError {
	ret := _m.Called(customerId, cardId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, string) *model.AppError); ok {
		r0 = rf(customerId, cardId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// MockPaymentRepository_SetDefaultCreditCard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetDefaultCreditCard'
type MockPaymentRepository_SetDefaultCreditCard_Call struct {
	*mock.Call
}

// SetDefaultCreditCard is a helper method to define mock.On call
//   - customerId string
//   - cardId string
func (_e *MockPaymentRepository_Expecter) SetDefaultCreditCard(customerId interface{}, cardId interface{}) *MockPaymentRepository_SetDefaultCreditCard_Call {
	return &MockPaymentRepository_SetDefaultCreditCard_Call{Call: _e.mock.On("SetDefaultCreditCard", customerId, cardId)}
}

func (_c *MockPaymentRepository_SetDefaultCreditCard_Call) Run(run func(customerId string, cardId string)) *MockPaymentRepository_SetDefaultCreditCard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockPaymentRepository_SetDefaultCreditCard_Call) Return(_a0 *model.AppError) *MockPaymentRepository_SetDefaultCreditCard_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPaymentRepository_SetDefaultCreditCard_Call) RunAndReturn(run func(string, string) *model.AppError) *MockPaymentRepository_SetDefaultCreditCard_Call {
	_c.Call.Return(run)
	return _c
}

// Subscribe provides a mock function with given fields: customerId, priceId
func (_m *MockPaymentRepository) Subscribe(customerId string, priceId string) (*model.Subscription, *model.AppError) {
	ret := _m.Called(customerId, priceId)

	var r0 *model.Subscription
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string, string) (*model.Subscription, *model.AppError)); ok {
		return rf(customerId, priceId)
	}
	if rf, ok := ret.Get(0).(func(string, string) *model.Subscription); ok {
		r0 = rf(customerId, priceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) *model.AppError); ok {
		r1 = rf(customerId, priceId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// MockPaymentRepository_Subscribe_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Subscribe'
type MockPaymentRepository_Subscribe_Call struct {
	*mock.Call
}

// Subscribe is a helper method to define mock.On call
//   - customerId string
//   - priceId string
func (_e *MockPaymentRepository_Expecter) Subscribe(customerId interface{}, priceId interface{}) *MockPaymentRepository_Subscribe_Call {
	return &MockPaymentRepository_Subscribe_Call{Call: _e.mock.On("Subscribe", customerId, priceId)}
}

func (_c *MockPaymentRepository_Subscribe_Call) Run(run func(customerId string, priceId string)) *MockPaymentRepository_Subscribe_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockPaymentRepository_Subscribe_Call) Return(_a0 *model.Subscription, _a1 *model.AppError) *MockPaymentRepository_Subscribe_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPaymentRepository_Subscribe_Call) RunAndReturn(run func(string, string) (*model.Subscription, *model.AppError)) *MockPaymentRepository_Subscribe_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPaymentRepository creates a new instance of MockPaymentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPaymentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPaymentRepository {
	mock := &MockPaymentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
