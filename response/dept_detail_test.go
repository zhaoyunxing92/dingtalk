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

package response

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeptDetail(t *testing.T) {
	str := `{
    "errcode": 0,
    "errmsg": "ok",
    "result": {
        "auto_add_user": true,
        "auto_approve_apply": false,
        "brief": "",
        "create_dept_group": true,
        "dept_group_chat_id": "chatcd94de28f11deae0dd34cac41684eb73",
        "dept_id": 560935057,
        "dept_manager_userid_list": [
            "16403377132689900"
        ],
        "dept_permits": [
            
        ],
        "group_contain_sub_dept": false,
        "hide_dept": true,
        "name": "golang",
        "order": 1,
        "org_dept_owner": "16403377132689900",
        "outer_dept": true,
        "outer_permit_depts": [
            554656655
        ],
        "outer_permit_users": [
            
        ],
        "parent_id": 1,
        "user_permits": [
            "manager164",
            "184919295227658120"
        ]
    },
    "request_id": "8p1f0ex1gw3y"
}`
	dept := &DeptDetail{}

	err := json.Unmarshal([]byte(str), dept)

	assert.Nil(t, err)
	assert.Equal(t, dept.Id, 560935057)
	assert.Equal(t, dept.Name, "golang")
	assert.Equal(t, dept.ParentId, 1)
	assert.Equal(t, dept.SourceIdentifier, "")
	assert.Equal(t, dept.CreateDeptGroup, true)
	assert.Equal(t, dept.AutoAddUser, true)
	assert.Equal(t, dept.OuterDept, true)
	assert.Equal(t, dept.HideDept, true)
	assert.Equal(t, dept.AutoApproveApply, false)
	assert.Equal(t, dept.FromUnionOrg, false)
	assert.Equal(t, dept.GroupContainSubDept, false)
	assert.Equal(t, dept.Order, 1)
	assert.Equal(t, dept.DeptGroupChatId, "chatcd94de28f11deae0dd34cac41684eb73")
	assert.Equal(t, dept.OrgDeptOwner, "16403377132689900")
	assert.Equal(t, dept.DeptManagerUseridList[0], "16403377132689900")
	assert.Equal(t, dept.UserPermitsDeptIds[0], 554656655)
	assert.Equal(t, dept.UserPermits[0], "manager164")
	assert.Equal(t, dept.UserPermits[1], "184919295227658120")
	assert.Equal(t, len(dept.UserPermitsUsers), 0)
	assert.Equal(t, len(dept.DeptPermits), 0)
}
