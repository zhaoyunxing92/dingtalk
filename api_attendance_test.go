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
	"encoding/json"
	"testing"

	"github.com/zhaoyunxing92/dingtalk/v2/constant/attendance"
	"github.com/zhaoyunxing92/dingtalk/v2/request"

	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

func TestDingTalk_GetAttendanceGroups(t *testing.T) {
	groups, err := client.GetAttendanceGroups(0, 10)

	assert.Nil(t, err)
	assert.NotNil(t, groups)
}

func TestGetAttendanceGroups_Json(t *testing.T) {
	str := `{
    "errcode": 0,
    "result": {
        "groups": [
            {
                "classes_list": [
                    "周一、二、三、四、五 A 09:30-18:30 ",
                    "周六、日 休息"
                ],
                "default_class_id": 686275261,
                "disable_check_when_rest": false,
                "disable_check_without_schedule": false,
                "enable_emp_select_class": true,
                "group_id": 792440025,
                "group_name": "技术考勤",
                "is_default": false,
                "manager_list": [
                    "manager164"
                ],
                "member_count": 1,
                "owner_user_id": "manager164",
                "selected_class": [
                    {
                        "class_id": 686275261,
                        "class_name": "A",
                        "sections": [
                            {
                                "times": [
                                    {
                                        "across": 0,
                                        "check_time": "1970-01-01 09:30:00",
                                        "check_type": "OnDuty"
                                    },
                                    {
                                        "across": 0,
                                        "check_time": "1970-01-01 18:30:00",
                                        "check_type": "OffDuty"
                                    }
                                ]
                            }
                        ],
                        "setting": {
                            "absenteeism_late_minutes": -1,
                            "class_setting_id": 599750584,
                            "is_off_duty_free_check": "N",
                            "permit_late_minutes": -1,
                            "serious_late_minutes": -1,
                            "work_time_minutes": 540
                        }
                    }
                ],
                "type": "FIXED",
                "work_day_list": [
                    "0",
                    "686275261",
                    "686275261",
                    "686275261",
                    "686275261",
                    "686275261",
                    "0"
                ]
            }
        ],
        "has_more": false
    },
    "request_id": "5wut35udt0h9"
}`
	group := &response.GetAttendanceGroup{}

	err := json.Unmarshal([]byte(str), group)

	assert.Nil(t, err)
	assert.Equal(t, group.RequestId, "5wut35udt0h9")
	assert.Equal(t, group.HasMore, false)
	assert.Equal(t, group.Code, 0)

	assert.Equal(t, group.Groups[0].GroupId, 792440025)
	assert.Equal(t, group.Groups[0].GroupName, "技术考勤")
	assert.Equal(t, group.Groups[0].Default, false)
	assert.Equal(t, group.Groups[0].EnableEmpSelectClass, true)
	assert.Equal(t, group.Groups[0].DisableCheckWhenRest, false)
	assert.Equal(t, group.Groups[0].DefaultClassId, 686275261)
	assert.Equal(t, group.Groups[0].DisableCheckWithoutSchedule, false)
	assert.Equal(t, group.Groups[0].MemberCount, 1)
	assert.Equal(t, group.Groups[0].ManagerList[0], "manager164")
	assert.Equal(t, group.Groups[0].OwnerUserId, "manager164")
	assert.Equal(t, group.Groups[0].Type, "FIXED")

	assert.Equal(t, len(group.Groups[0].WorkDayList), 7)

	assert.Equal(t, group.Groups[0].SelectedClass[0].ClassId, 686275261)
	assert.Equal(t, group.Groups[0].SelectedClass[0].ClassName, "A")
	assert.Equal(t, group.Groups[0].SelectedClass[0].Setting.ClassSettingId, 599750584)
	assert.Equal(t, group.Groups[0].SelectedClass[0].Setting.AbsenteeismLateMinutes, -1)
	assert.Equal(t, group.Groups[0].SelectedClass[0].Setting.IsOffDutyFreeCheck, "N")
	assert.Equal(t, group.Groups[0].SelectedClass[0].Setting.PermitLateMinutes, -1)
	assert.Equal(t, group.Groups[0].SelectedClass[0].Setting.SeriousLateMinutes, -1)
	assert.Equal(t, group.Groups[0].SelectedClass[0].Setting.WorkTimeMinutes, 540)

	assert.Equal(t, group.Groups[0].SelectedClass[0].Sections[0].Times[0].CheckTime, "1970-01-01 09:30:00")
	assert.Equal(t, group.Groups[0].SelectedClass[0].Sections[0].Times[0].CheckType, "OnDuty")
	assert.Equal(t, group.Groups[0].SelectedClass[0].Sections[0].Times[0].Across, 0)

	assert.Equal(t, group.Groups[0].SelectedClass[0].Sections[0].Times[1].CheckTime, "1970-01-01 18:30:00")
	assert.Equal(t, group.Groups[0].SelectedClass[0].Sections[0].Times[1].CheckType, "OffDuty")
	assert.Equal(t, group.Groups[0].SelectedClass[0].Sections[0].Times[1].Across, 0)
}

func TestGetAttendanceUserGroup(t *testing.T) {
	group, err := client.GetAttendanceUserGroup("manager164")

	assert.Nil(t, err)
	assert.NotNil(t, group)
	assert.Equal(t, group.AttendanceUserGroup.Type, "FIXED")
	assert.Equal(t, group.AttendanceUserGroup.Name, "技术考勤")
	assert.Equal(t, group.AttendanceUserGroup.GroupId, 792440025)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Name, "考勤班次")
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].ClassId, 686275261)

	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Setting.RestEndTime.Across, 0)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Setting.RestEndTime.CheckType, "OffDuty")
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Setting.RestEndTime.CheckTime, "1970-01-01 13:00:00")

	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Setting.RestBeginTime.Across, 0)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Setting.RestBeginTime.CheckType, "OnDuty")
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Setting.RestBeginTime.CheckTime, "1970-01-01 12:00:00")

	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[0].Across, 0)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[0].BeginMin, 1)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[0].EndMin, 1)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[0].CheckType, "OnDuty")
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[0].CheckTime, "1970-01-01 09:30:00")

	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[1].Across, 0)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[1].BeginMin, 1)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[1].EndMin, 1)
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[1].CheckType, "OffDuty")
	assert.Equal(t, group.AttendanceUserGroup.Classes[0].Sections[0].Times[1].CheckTime, "1970-01-01 18:30:00")
}

func TestDingTalk_GetAttendanceGroupMinimalism(t *testing.T) {
	minimalism, err := client.GetAttendanceGroupMinimalism("manager164", 0)

	assert.Nil(t, err)
	assert.Equal(t, minimalism.MinimalismGroup.Attendance[0].Id, 792440025)
	assert.Equal(t, minimalism.MinimalismGroup.Attendance[0].Name, "技术考勤")
}

func TestDingTalk_GetAttendanceGroupDetail(t *testing.T) {
	detail, err := client.GetAttendanceGroupDetail("manager164", 792440025)

	assert.Nil(t, err)
	assert.Equal(t, detail.AttendanceGroup.Id, 792440025)
	assert.Equal(t, detail.AttendanceGroup.Name, "技术考勤")
	assert.Equal(t, detail.AttendanceGroup.AddressList[0], "绿城未来park")
	assert.Equal(t, detail.AttendanceGroup.MemberCount, 3)
	assert.Equal(t, detail.AttendanceGroup.OwnerUserId, "manager164")
	assert.Equal(t, detail.AttendanceGroup.ShiftIds[0], 686275261)
	assert.Equal(t, detail.AttendanceGroup.Type, "FIXED")
	assert.Equal(t, detail.AttendanceGroup.Wifis[0], "wifi")
	assert.Equal(t, len(detail.AttendanceGroup.WorkDayList), 7)
}

func TestDingTalk_SearchAttendanceGroup(t *testing.T) {
	group, err := client.SearchAttendanceGroup("manager164", "技术")

	assert.Nil(t, err)
	assert.Equal(t, len(group.AttendanceGroups), 1)
}

func TestDingTalk_CreateAttendanceGroup_TURN(t *testing.T) {
	t.Skip()
	group := request.NewCreateAttendanceGroup("manager164",
		"创建排班制考勤组",
		attendance.TURN,
		[]request.AttendanceMember{request.NewAttendanceMember("manager164", "")}).
		Build()

	res, err := client.CreateAttendanceGroup(group)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res.AttendanceGroup.Name, "创建排班制考勤组")
}
