/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package request

import (
	"encoding/json"
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestNewInactiveUser(t *testing.T) {
	iu := NewInactiveUser(false, 0, 10, "20211104").
		SetDeptIds(123, 123, 456, 789, 123).
		Build()

	assert.Equal(t, len(iu.DeptIds), 3)
	js, err := json.Marshal(iu)
	assert.Nil(t, err)
	t.Log(string(js))

}
