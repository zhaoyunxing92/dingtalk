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

	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/employee"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
)

func TestDingTalk_GetHrmEmployee(t *testing.T) {
	user, err := client.GetHrmEmployee(1, 10, []employee.Status{employee.Formal, employee.Unknown, employee.ProbationPeriod})

	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestDingTalk_GetHrmToBeHiredEmployee(t *testing.T) {
	user, err := client.GetHrmToBeHiredEmployee(1, 10)

	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestDingTalk_GetHrmResignEmployeeIds(t *testing.T) {
	user, err := client.GetHrmResignEmployeeIds(0, 50)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, len(user.UserIds), 32)
}

func TestDingTalk_GetHrmResignEmployee(t *testing.T) {
	user, err := client.GetHrmResignEmployee([]string{"223528591438654612", "055354431324052749"})

	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestDingTalk_HrmCreateEmployee(t *testing.T) {
	emp := request.NewHrmCreateEmployee("老刘", "18867870721").Build()

	user, err := client.HrmCreateEmployee(emp)

	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestDingTalk_GetHrmEmployeeField(t *testing.T) {
	res, err := client.GetHrmEmployeeField(1244553273, []string{"manager164"}, []string{})

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_UpdateHrmEmployeeField(t *testing.T) {
	value := request.EmpFieldValue{FieldCode: "sys00-remark", Value: "备注"}
	field := request.EmpField{Values: []request.EmpFieldValue{value}}
	group := request.UpdateHrmEmpFieldGroup{GroupId: "sys00", Fields: []request.EmpField{field}}

	req := &request.UpdateHrmEmpField{
		AgentId: 1244553273,
		Param: request.UpdateHrmEmpFieldParam{
			UserId:        "manager164",
			EmpFieldGroup: []request.UpdateHrmEmpFieldGroup{group},
		},
	}

	res, err := client.UpdateHrmEmployeeField(req)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetHrmMeta(t *testing.T) {
	res, err := client.GetHrmMeta(1244553273)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
