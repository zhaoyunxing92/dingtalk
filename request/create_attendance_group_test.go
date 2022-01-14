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

	"github.com/zhaoyunxing92/dingtalk/v2/constant/attendance"
)

func TestNewCreateAttendanceGroup(t *testing.T) {
	group := NewCreateAttendanceGroup("manager164",
		"创建固定班制考勤组",
		attendance.FIXED,
		[]AttendanceMember{NewAttendanceMember("manager164", "")}).
		SetWorkdayClassList([]string{"0", "12", "12", "12", "12", "12", "0"}).
		SetOpUserId("userId").
		SetCorpId("corpId").
		SetEnableEmpSelectClass(true).
		SetEnableNextDay(true).
		SetEnableOutsideCheck(true).
		SetDefaultClassId(1234).
		SetModifyMember(true).
		SetCheckNeedHealthyCode(true).
		SetEnableFaceCheck(true).
		SetDisableCheckWithoutSchedule(true).
		SetDisableCheckWhenRest(true).
		SetAttendanceType(attendance.FIXED).
		SetEnableCameraCheck(true).
		SetEnablePositionBle(true).
		SetEnableOutsideCameraCheck(true).
		SetName("name").
		SetBleDevices(12345).
		SetWorkdayClassList([]string{"0", "12", "12", "12", "12", "12", "0"}).
		SetResourcePermissions("schedule", "outSideCheck", "groupType",
			"groupMember", "checkPositionType", "cameraCheck",
			"checkTime", "overTimeRule").
		SetPositions("address", "latitude", "longitude", "accuracy", "title", "corpId").
		SetOffset(12).
		SetFreeCheckDayStartMinOffset(12).
		SetSkipHolidays(true).
		SetOwner("owner").
		SetAttendanceShifts(123).
		SetManagers("1234").
		SetSpecialDays("SetSpecialDays").
		SetMembers("userId", "corpId").
		SetAttendanceWifi("addr", "ssId", "corpId").
		Build()

	t.Log(group.String())
}
