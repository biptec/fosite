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
// Source: github.com/ory/fosite/handler/oauth2 (interfaces: AuthorizeCodeStorage)

package internal

import (
	context "context"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// Mock of AuthorizeCodeStorage interface
type MockAuthorizeCodeStorage struct {
	ctrl     *gomock.Controller
	recorder *_MockAuthorizeCodeStorageRecorder
}

// Recorder for MockAuthorizeCodeStorage (not exported)
type _MockAuthorizeCodeStorageRecorder struct {
	mock *MockAuthorizeCodeStorage
}

func NewMockAuthorizeCodeStorage(ctrl *gomock.Controller) *MockAuthorizeCodeStorage {
	mock := &MockAuthorizeCodeStorage{ctrl: ctrl}
	mock.recorder = &_MockAuthorizeCodeStorageRecorder{mock}
	return mock
}

func (_m *MockAuthorizeCodeStorage) EXPECT() *_MockAuthorizeCodeStorageRecorder {
	return _m.recorder
}

func (_m *MockAuthorizeCodeStorage) CreateAuthorizeCodeSession(_param0 context.Context, _param1 string, _param2 fosite.Requester) error {
	ret := _m.ctrl.Call(_m, "CreateAuthorizeCodeSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAuthorizeCodeStorageRecorder) CreateAuthorizeCodeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAuthorizeCodeSession", arg0, arg1, arg2)
}

func (_m *MockAuthorizeCodeStorage) DeleteAuthorizeCodeSession(_param0 context.Context, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteAuthorizeCodeSession", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockAuthorizeCodeStorageRecorder) DeleteAuthorizeCodeSession(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAuthorizeCodeSession", arg0, arg1)
}

func (_m *MockAuthorizeCodeStorage) GetAuthorizeCodeSession(_param0 context.Context, _param1 string, _param2 fosite.Session) (fosite.Requester, error) {
	ret := _m.ctrl.Call(_m, "GetAuthorizeCodeSession", _param0, _param1, _param2)
	ret0, _ := ret[0].(fosite.Requester)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockAuthorizeCodeStorageRecorder) GetAuthorizeCodeSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAuthorizeCodeSession", arg0, arg1, arg2)
}
