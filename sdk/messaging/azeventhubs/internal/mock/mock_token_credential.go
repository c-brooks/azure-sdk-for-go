// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// Code generated by MockGen. DO NOT EDIT.
// Source: ../../../azcore/internal/exported/exported.go

package mock

import (
	context "context"
	reflect "reflect"

	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/policy"
	gomock "github.com/golang/mock/gomock"
)

// MockTokenCredential is a mock of TokenCredential interface.
type MockTokenCredential struct {
	ctrl     *gomock.Controller
	recorder *MockTokenCredentialMockRecorder
}

// MockTokenCredentialMockRecorder is the mock recorder for MockTokenCredential.
type MockTokenCredentialMockRecorder struct {
	mock *MockTokenCredential
}

// NewMockTokenCredential creates a new mock instance.
func NewMockTokenCredential(ctrl *gomock.Controller) *MockTokenCredential {
	mock := &MockTokenCredential{ctrl: ctrl}
	mock.recorder = &MockTokenCredentialMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenCredential) EXPECT() *MockTokenCredentialMockRecorder {
	return m.recorder
}

// GetToken mocks base method.
func (m *MockTokenCredential) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", ctx, options)
	ret0, _ := ret[0].(azcore.AccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockTokenCredentialMockRecorder) GetToken(ctx, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockTokenCredential)(nil).GetToken), ctx, options)
}
