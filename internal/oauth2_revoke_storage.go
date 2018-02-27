/*
 * Copyright © 2017-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @Copyright 	2017-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 *
 */

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: TokenRevocationStorage)

package internal

import (
	context "context"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// Mock of TokenRevocationStorage interface
type MockTokenRevocationStorage struct {
	ctrl     *gomock.Controller
	recorder *_MockTokenRevocationStorageRecorder
}

// Recorder for MockTokenRevocationStorage (not exported)
type _MockTokenRevocationStorageRecorder struct {
	mock *MockTokenRevocationStorage
}

func NewMockTokenRevocationStorage(ctrl *gomock.Controller) *MockTokenRevocationStorage {
	mock := &MockTokenRevocationStorage{ctrl: ctrl}
	mock.recorder = &_MockTokenRevocationStorageRecorder{mock}
	return mock
}

func (_m *MockTokenRevocationStorage) EXPECT() *_MockTokenRevocationStorageRecorder {
	return _m.recorder
}

func (_m *MockTokenRevocationStorage) CreateAccessTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "CreateAccessTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTokenRevocationStorageRecorder) CreateAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAccessTokenSession", arg0, arg1, arg2)
}

func (_m *MockTokenRevocationStorage) CreateRefreshTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "CreateRefreshTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTokenRevocationStorageRecorder) CreateRefreshTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateRefreshTokenSession", arg0, arg1, arg2)
}

func (_m *MockTokenRevocationStorage) DeleteAccessTokenSession(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteAccessTokenSession", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTokenRevocationStorageRecorder) DeleteAccessTokenSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAccessTokenSession", arg0, arg1)
}

func (_m *MockTokenRevocationStorage) DeleteRefreshTokenSession(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteRefreshTokenSession", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTokenRevocationStorageRecorder) DeleteRefreshTokenSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteRefreshTokenSession", arg0, arg1)
}

func (_m *MockTokenRevocationStorage) GetAccessTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Session) (fosite.Requester, error) {
	ret := _m.ctrl.Call(_m, "GetAccessTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTokenRevocationStorageRecorder) GetAccessTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAccessTokenSession", arg0, arg1, arg2)
}

func (_m *MockTokenRevocationStorage) GetRefreshTokenSession(_param0 context.Context, _param1 string, _param2 fosite.Session) (fosite.Requester, error) {
	ret := _m.ctrl.Call(_m, "GetRefreshTokenSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTokenRevocationStorageRecorder) GetRefreshTokenSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRefreshTokenSession", arg0, arg1, arg2)
}

func (_m *MockTokenRevocationStorage) RevokeAccessToken(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "RevokeAccessToken", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTokenRevocationStorageRecorder) RevokeAccessToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RevokeAccessToken", arg0, arg1)
}

func (_m *MockTokenRevocationStorage) RevokeRefreshToken(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "RevokeRefreshToken", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockTokenRevocationStorageRecorder) RevokeRefreshToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RevokeRefreshToken", arg0, arg1)
}
