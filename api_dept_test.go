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

	"github.com/zhaoyunxing92/dingtalk/v2/constant/language"

	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
)

func TestDingTalk_GetDeptSimpleUserInfo(t *testing.T) {
	info, err := client.GetDeptSimpleUserInfo(
		request.NewDeptSimpleUserInfo(1, 0, 1).
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, info.Result.DeptUsers)
}

func TestDingTalk_GetDeptUserIds(t *testing.T) {
	userId, err := client.GetDeptUserIds(request.NewDeptUserId(1))

	assert.Nil(t, err)
	assert.NotNil(t, userId.Result.UserIds)
}

func TestDingTalk_GetDeptDetailUserInfo(t *testing.T) {
	res, err := client.GetDeptDetailUserInfo(
		request.NewDeptDetailUserInfo(1, 0, 10).
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res.Page.List)
}

func TestDingTalk_CreateDept(t *testing.T) {
	t.Skip()
	res, err := client.CreateDept(
		request.NewCreateDept("test", 1).
			SetHideDept(false).
			SetDeptPermits(1, 1, 2).
			SetOrder(1).
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_DeleteDept(t *testing.T) {
	t.Skip()
	res, err := client.DeleteDept(560900478)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetDeptDetail(t *testing.T) {
	res, err := client.GetDeptDetail(request.NewDeptDetail(1).Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.Detail.Name, "钉钉小程序开发团队")
}

func TestDingTalk_UpdateDept(t *testing.T) {
	res, err := client.UpdateDept(
		request.NewUpdateDept(560935057).
			SetHideDept(true).
			SetOuterDept(true).
			SetDeptPermits(1, 554656655).
			SetUserPermits("manager164").
			SetUserPermitsDeptIds(1, 554656655).
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetDeptList(t *testing.T) {
	res, err := client.GetDeptList(
		request.NewDeptList().
			SetDeptId(1).
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetSubDeptList(t *testing.T) {
	res, err := client.GetSubDeptList(1)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetParentIdsByDeptId(t *testing.T) {
	res, err := client.GetParentIdsByDeptId(554656655)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res.Result.ParentIds)
}

func TestDingTalk_GetParentIdsByUserId(t *testing.T) {
	res, err := client.GetParentIdsByUserId("manager164")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res.Result.Parents)
}

func TestDingTalk_FetchDeptList(t *testing.T) {
	res, err := client.FetchDeptList(1, true, language.ZH_CN)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
