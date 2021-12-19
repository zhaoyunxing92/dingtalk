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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/employee"
)

func TestNewGetHrmEmployee(t *testing.T) {
	e := NewGetHrmEmployee(1, 0,
		[]employee.Status{employee.Formal, employee.Formal, employee.ToBeResigned, employee.Unknown})

	assert.NotNil(t, e)
	assert.Equal(t, e.Offset, 1)
	assert.Equal(t, e.Size, 0)
	assert.Equal(t, e.Status, "3,5,-1")
}
