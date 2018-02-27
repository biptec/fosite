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

package fosite_test

import (
	"testing"

	. "github.com/ory/fosite"
	"github.com/stretchr/testify/assert"
)

func TestAccessResponse(t *testing.T) {
	ar := NewAccessResponse()
	ar.SetAccessToken("access")
	ar.SetTokenType("bearer")
	ar.SetExtra("access_token", "invalid")
	ar.SetExtra("foo", "bar")
	assert.Equal(t, "access", ar.GetAccessToken())
	assert.Equal(t, "bearer", ar.GetTokenType())
	assert.Equal(t, "bar", ar.GetExtra("foo"))
	assert.Equal(t, map[string]interface{}{
		"access_token": "access",
		"token_type":   "bearer",
		"foo":          "bar",
	}, ar.ToMap())
}
