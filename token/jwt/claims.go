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

package jwt

import "time"

// Mapper is the interface used internally to map key-value pairs
type Mapper interface {
	ToMap() map[string]interface{}
	Add(key string, value interface{})
	Get(key string) interface{}
}

// ToString will return a string representation of a map
func ToString(i interface{}) string {
	if i == nil {
		return ""
	}

	if s, ok := i.(string); ok {
		return s
	}

	if sl, ok := i.([]string); ok {
		if len(sl) == 1 {
			return sl[0]
		}
	}

	return ""
}

// ToTime will try to convert a given input to a time.Time structure
func ToTime(i interface{}) time.Time {
	if i == nil {
		return time.Time{}
	}

	if t, ok := i.(int64); ok {
		return time.Unix(t, 0).UTC()
	} else if t, ok := i.(float64); ok {
		return time.Unix(int64(t), 0).UTC()
	} else if t, ok := i.(time.Time); ok {
		return t
	}

	return time.Time{}
}

// Filter will filter out elemets based on keys in a given input map na key-slice
func Filter(elements map[string]interface{}, keys ...string) map[string]interface{} {
	var keyIdx = make(map[string]bool)
	var result = make(map[string]interface{})

	for _, key := range keys {
		keyIdx[key] = true
	}

	for k, e := range elements {
		if _, ok := keyIdx[k]; !ok {
			result[k] = e
		}
	}

	return result
}

// Copy will copy all elements in a map and return a new representational map
func Copy(elements map[string]interface{}) (result map[string]interface{}) {
	result = make(map[string]interface{}, len(elements))
	for k, v := range elements {
		result[k] = v
	}

	return result
}
