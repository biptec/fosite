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
// Source: github.com/ory/fosite/handler/openid (interfaces: OpenIDConnectTokenStrategy)

package internal

import (
	context "context"

	gomock "github.com/golang/mock/gomock"
	fosite "github.com/ory/fosite"
)

// Mock of OpenIDConnectTokenStrategy interface
type MockOpenIDConnectTokenStrategy struct {
	ctrl     *gomock.Controller
	recorder *_MockOpenIDConnectTokenStrategyRecorder
}

// Recorder for MockOpenIDConnectTokenStrategy (not exported)
type _MockOpenIDConnectTokenStrategyRecorder struct {
	mock *MockOpenIDConnectTokenStrategy
}

func NewMockOpenIDConnectTokenStrategy(ctrl *gomock.Controller) *MockOpenIDConnectTokenStrategy {
	mock := &MockOpenIDConnectTokenStrategy{ctrl: ctrl}
	mock.recorder = &_MockOpenIDConnectTokenStrategyRecorder{mock}
	return mock
}

func (_m *MockOpenIDConnectTokenStrategy) EXPECT() *_MockOpenIDConnectTokenStrategyRecorder {
	return _m.recorder
}

func (_m *MockOpenIDConnectTokenStrategy) GenerateIDToken(_param0 context.Context, _param1 fosite.Requester) (string, error) {
	ret := _m.ctrl.Call(_m, "GenerateIDToken", _param0, _param1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockOpenIDConnectTokenStrategyRecorder) GenerateIDToken(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GenerateIDToken", arg0, arg1)
}
