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
)

func TestNewUpdateDept(t *testing.T) {
	str := NewUpdateDept(0).
		SetHideDept(true).
		SetOuterDept(true).
		SetUserPermits("123", "456", "").
		SetOrder(0).
		SetParentId(0).
		SetName("name").
		SetCreateDeptGroup(false).
		SetAutoApproveApply(false).
		SetAutoAddUser(false).
		SetSourceIdentifier("id").
		SetDeptManagerUseridList("id", "1234").
		SetUserPermitsUserIds("123", "123", "", "456").
		SetUserPermitsDeptIds(1, 0, 345, 567, 345).
		SetDeptPermits(1, 0, 345, 567, 345).
		Build()

	assert.NotNil(t, str)
	assert.Equal(t, len(str.UserPermitsUserIds), 2)
	assert.Equal(t, len(str.UserPermitsDeptIds), 3)
	assert.Equal(t, len(str.UserPermit), 2)
	assert.Equal(t, len(str.DeptPermit), 3)
	t.Log(str.String())
}

func TestNewEntryUserPermitsUpdateDept(t *testing.T) {
	str := NewUpdateDept(0).
		SetHideDept(true).
		SetOuterDept(true).
		// SetUserPermits("123", "456", "").
		SetOrder(0).
		SetUserPermitsUserIds("123", "123", "", "456").
		SetUserPermitsDeptIds(1, 0, 345, 567, 345).
		SetDeptPermits(1, 0, 345, 567, 345).
		Build()

	assert.NotNil(t, str)
	assert.Equal(t, len(str.UserPermitsUserIds), 2)
	assert.Equal(t, len(str.UserPermitsDeptIds), 3)
	// assert.Equal(t, len(str.UserPermit), 2)
	assert.Equal(t, len(str.DeptPermit), 3)
	t.Log(str.String())
}

func TestNewEntryUserPermitsUserIdsUpdateDept(t *testing.T) {
	str := NewUpdateDept(0).
		SetHideDept(true).
		SetOuterDept(true).
		// SetUserPermits("123", "456", "").
		SetOrder(0).
		// SetUserPermitsUserIds("123", "123", "", "456").
		SetUserPermitsDeptIds(1, 0, 345, 567, 345).
		SetDeptPermits(1, 0, 345, 567, 345).
		Build()

	assert.NotNil(t, str)
	// assert.Equal(t, len(str.UserPermitsUserIds), 2)
	assert.Equal(t, len(str.UserPermitsDeptIds), 3)
	// assert.Equal(t, len(str.UserPermit), 2)
	assert.Equal(t, len(str.DeptPermit), 3)
	t.Log(str.String())
}

func TestNewHideDeptUpdateDept(t *testing.T) {
	str := NewUpdateDept(0).
		SetHideDept(false).
		SetDeptPermits(1, 1, 2).
		SetOuterDeptOnlySelf(true).
		SetOrder(1).
		Build()

	assert.NotNil(t, str)
	assert.Equal(t, len(str.DeptPermit), 2)
	assert.Nil(t, str.OuterDeptOnlySelf)
	t.Log(str.String())
}
