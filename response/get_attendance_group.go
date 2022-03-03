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

type GetAttendanceGroup struct {
	Response

	AttendanceGroups `json:"result"`
}

type AttendanceGroups struct {
	HasMore bool `json:"has_more"`

	Groups []AttendanceGroup `json:"groups"`
}

type AttendanceGroup struct {
	// 考勤组ID
	GroupId int `json:"group_id"`

	//是否是默认考勤组。
	//
	//false：不是
	//
	//true：是
	Default bool `json:"is_default"`

	// 考勤组名称
	GroupName string `json:"group_name"`

	//考勤类型。
	//
	//FIXED为固定排班
	//
	//TURN为轮班排班
	//
	//NONE为无班次
	Type string `json:"type"`

	// 考勤组对应的考勤班次列表
	SelectedClass []SelectedClass `json:"selected_class"`

	// 一周的班次时间展示列表
	ClassesList []string `json:"classes_list"`

	DisableCheckWhenRest bool `json:"disable_check_when_rest"`

	DisableCheckWithoutSchedule bool `json:"disable_check_without_schedule"`

	EnableEmpSelectClass bool `json:"enable_emp_select_class"`

	// 考勤组主负责人
	OwnerUserId string `json:"owner_user_id"`

	//固定班次的工作日班次。
	//
	//说明
	//0表示休息。
	//数组内的值，从左到右依次代表周日到周六，每日的排班情况。
	WorkDayList []string `json:"work_day_list"`

	// 默认班次ID
	DefaultClassId int `json:"default_class_id"`

	// 考勤组子负责人
	ManagerList []string `json:"manager_list"`

	// 参与考勤人员人数。
	MemberCount int `json:"member_count"`
}

// SelectedClass 考勤组对应的考勤班次列表
type SelectedClass struct {

	// 考勤班次ID
	ClassId int `json:"class_id"`

	// 考勤班次名称
	ClassName string `json:"class_name"`

	// 未排班时是否允许员工选择班次打卡。
	EnableEmpSelectClass bool `json:"enable_emp_select_class"`

	// 未排班时是否禁止员工打卡
	DisableCheckWithoutSchedule bool `json:"disable_check_without_schedule"`

	//休息日打卡是否需审批：
	//
	//true：需要
	//
	//false：不需要
	DisableCheckWhenRest bool `json:"disable_check_when_rest"`

	// 考勤组班次配置
	Setting `json:"setting"`

	// 班次打卡时间段
	Sections []Section `json:"sections"`
}

// Setting 考勤组班次配置
type Setting struct {

	// 考勤组班次ID
	ClassSettingId int `json:"class_setting_id"`

	// 旷工迟到时长，单位分钟（-1表示关闭该功能）
	AbsenteeismLateMinutes int `json:"absenteeism_late_minutes"`

	//是否强制下班打卡。
	//
	//Y：不强制打卡
	//
	//N：强制打卡
	IsOffDutyFreeCheck string `json:"is_off_duty_free_check"`

	// 休息开始时间，只有一个时间段的班次有
	RestBeginTime `json:"rest_begin_time"`

	// 允许迟到时长，单位分钟（-1表示关闭该功能）
	PermitLateMinutes int `json:"permit_late_minutes"`

	// 严重迟到时长，单位分钟（-1表示关闭该功能）
	SeriousLateMinutes int `json:"serious_late_minutes"`

	// 工作时长，单位分钟（-1表示关闭该功能）
	WorkTimeMinutes int `json:"work_time_minutes"`

	// 休息结束时间，只有一个时间段的班次有
	RestEndTime `json:"rest_end_time"`
}

// Section 班次打卡时间段
type Section struct {

	// 时间段列表。
	Times []Time `json:"times"`
}

// Time 时间段列表。
type Time struct {
	// 打卡时间跨度
	Across int `json:"across"`

	// 打卡时间。
	CheckTime string `json:"check_time"`

	//打卡类型。
	//
	//OnDuty：上班
	//
	//OffDuty：下班
	CheckType string `json:"check_type"`
}

// RestBeginTime 休息开始时间，只有一个时间段的班次有
type RestBeginTime struct {
	// 开始时间
	CheckTime string `json:"check_time"`
}

// RestEndTime 休息结束时间，只有一个时间段的班次有
type RestEndTime struct {
	// 结束时间
	CheckTime string `json:"check_time"`
}
