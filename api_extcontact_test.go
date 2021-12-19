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

package dingtalk

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/request"
)

func TestDingTalk_CreateExtContact(t *testing.T) {
	res, err := client.CreateExtContact(
		request.NewCreateExtContact("张三", "18058762219", "86", "manager164",
			1587784123, 1587784123, 1587784124, 456).
			SetTitle("资深开发").
			SetAddress("杭州").
			SetShareDept(123, 123, 678).
			SetRemark("技术人员").
			SetCompanyName("小番茄").
			SetShareUser("manager164").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_DeleteExtContact(t *testing.T) {
	res, err := client.DeleteExtContact("01131104492477488")

	assert.NotNil(t, err)
	assert.Equal(t, res.Ok(), false)
}

func TestDingTalk_UpdateExtContact(t *testing.T) {
	res, err := client.UpdateExtContact(
		request.NewUpdateExtContact("01491341662086", "李四", "manager164",
			1587784123, 1587784123, 1587784124, 1587784125).
			SetTitle("供应商").
			SetAddress("杭州").
			SetShareDept(560935057, 123, 678).
			SetRemark("技术人员").
			SetCompanyName("小番茄").
			SetShareUser("manager164").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetExtContact(t *testing.T) {
	res, err := client.GetExtContact(0, 10)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetExtContactLabel(t *testing.T) {
	res, err := client.GetExtContactLabel(0, 10)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(res.Results), 4)
}

func TestDingTalk_GetExtContactDetail(t *testing.T) {
	res, err := client.GetExtContactDetail("01491341662086")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.Name, "李四")
	assert.Equal(t, res.CompanyName, "小番茄")
	assert.Equal(t, res.Address, "杭州")
	assert.Equal(t, res.FollowerUser, "manager164")
	assert.Equal(t, len(res.Labels), 3)
	assert.Equal(t, len(res.ShareDept), 1)
	assert.Equal(t, res.Mobile, "18058762217")
}
