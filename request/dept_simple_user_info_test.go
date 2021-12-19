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
	"github.com/zhaoyunxing92/dingtalk/v2/constant/language"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/order"
)

func TestNewDeptSimpleUserInfo(t *testing.T) {
	dept := NewDeptSimpleUserInfo(1234, 0, 10).
		SetOrderField(order.EntryDesc).
		SetLanguage(language.EN_US).
		SetContainAccessLimit(true).
		Build()

	assert.NotNil(t, dept)
	assert.Equal(t, dept.Size, 10)
	assert.Equal(t, dept.Cursor, 0)
	assert.Equal(t, dept.DeptId, 1234)
	assert.Equal(t, dept.ContainAccessLimit, true)
	assert.Equal(t, dept.Language, string(language.EN_US))
	assert.Equal(t, dept.OrderField, string(order.EntryDesc))
}
