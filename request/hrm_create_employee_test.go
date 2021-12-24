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
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/employee"
)

func TestNewHrmCreateEmployee(t *testing.T) {
	emp := NewHrmCreateEmployee("老刘", "18357154439").
		SetEntryTime(time.Now()).
		SetOperator("operator").
		SetDeptList([]int{1, 560935057, 560935057}).
		SetMainDeptId(560935057).
		SetMainDeptName("deptName").
		SetPosition("position").
		SetWorkPlace("workPlace").
		SetJobNumber("jobNumber").
		SetEmployeeType(employee.FullTimeMold).
		Build()

	assert.NotNil(t, emp)
	assert.Equal(t, emp.Param.Name, "老刘")
	assert.Equal(t, emp.Param.Mobile, "18357154439")
	assert.Equal(t, emp.Param.Operator, "operator")
	assert.Equal(t, emp.Param.Ext.WorkPlace, "workPlace")
	assert.Equal(t, emp.Param.Ext.MainDeptName, "deptName")
	assert.Equal(t, emp.Param.Ext.EmployeeType, string(employee.FullTimeMold))

	t.Log(emp.String())
}

func TestNewHrmCreateEmployee_Simple(t *testing.T) {
	emp := NewHrmCreateEmployee("老刘", "18357154439").Build()

	assert.NotNil(t, emp)
	assert.Equal(t, emp.Param.Name, "老刘")
	assert.Equal(t, emp.Param.Mobile, "18357154439")

	t.Log(emp.String())
}
