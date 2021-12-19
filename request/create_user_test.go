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
)

import (
	"github.com/stretchr/testify/assert"
)

func TestNewCreateUser(t *testing.T) {
	user := NewCreateUser("zhangsan", "18357154439", 123, 123, 7890, 123456, 123456).
		SetDeptOrder(123, 1).
		SetDeptOrder(456, 2).
		SetName("lisi").
		SetMobile("110").
		SetHideMobile(false).
		SetTelephone("telephone").
		SetJobNumber("1234").
		SetTitle("开发").
		SetEmail("email").
		SetOrgEmail("org").
		SetOrgEmailType("type").
		SetRemark("牛逼开发").
		SetManagerUserId("managerUserId").
		SetWorkPlace("work").
		SetDeptTitle(123, "开发").
		SetDeptTitle(456, "测试").
		SetExtension("测试").
		SetSeniorMode(true).
		SetHiredDate(123456789).
		SetLoginEmail("login").
		SetExclusiveAccount(true).
		SetSetExclusiveAccountType("exclusive").
		SetLoginId("login").
		SetPassword("pwd").
		Build()

	assert.NotNil(t, user)
}
