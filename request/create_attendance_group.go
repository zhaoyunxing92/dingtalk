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
	"encoding/json"

	"github.com/zhaoyunxing92/dingtalk/v2/constant/attendance"
)

type CreateAttendanceGroup struct {
	// 操作人的userid
	OpUserId string `json:"op_user_id" validate:"required"`

	// 考勤组相关信息
	AttendanceGroup *AttendanceGroup `json:"top_group"`
}

type AttendanceGroup struct {

	// 考勤wifi打卡相关配置信息
	Wifis []AttendanceWifi `json:"wifis,omitempty"`

	//是否可以外勤打卡。
	//
	//true：允许（默认值）
	//
	//false：不允许
	EnableCameraCheck *bool `json:"enable_camera_check,omitempty"`

	//自由工时考勤组考勤开始时间与当天0点偏移分钟数。
	//
	//例如：540表示9:00
	FreeCheckDayStartMinOffset *int `json:"freecheck_day_start_min_offset,omitempty"`

	// 特殊日期配置 {"onDuty":{1400000:123,1400001:123},"offDuty":[1400000,1400001]}
	SpecialDays string `json:"special_days,omitempty"`

	//是否第二天生效。
	//
	//true：是
	//
	//false：否
	EnableNextDay *bool `json:"enable_next_day,omitempty"`

	// 班次相关配置信息
	AttendanceShifts []AttendanceShift `json:"shift_vo_list,omitempty"`

	//休息日打卡是否需审批：
	//
	//true：需要
	//
	//false：不需要
	DisableCheckWhenRest *bool `json:"disable_check_when_rest,omitempty"`

	//考考勤组类型：
	//
	//FIXED：固定班制考勤组
	//
	//TURN：排班制考勤组
	//
	//NONE：自由工时考勤组
	Type string `json:"type,omitempty" validate:"required,oneof=FIXED TURN NONE"`

	// 未排班时是否禁止员工打卡
	DisableCheckWithoutSchedule *bool `json:"disable_check_without_schedule,omitempty"`

	// 是否有修改考勤组成员相关信
	ModifyMember *bool `json:"modify_member,omitempty"`

	//打卡是否需要健康码：
	//
	//true：开启
	//
	//false：关闭（默认值）
	CheckNeedHealthyCode *bool `json:"check_need_healthy_code,omitempty"`

	// 考勤组成员相关设置信息
	Members []AttendanceMember `json:"members,omitempty" validate:"required"`

	// 是否启用蓝牙定位
	EnablePositionBle *bool `json:"enable_position_ble,omitempty"`

	// 考勤组ID
	Id string `json:"id,omitempty"`

	// 考勤组子管理员userid列表
	Managers []string `json:"manager_list,omitempty"`

	// 考勤组负责人
	Owner string `json:"owner,omitempty"`

	//是否跳过节假日。
	//
	//true：跳过（默认值）
	//
	//false：不跳过
	SkipHolidays *bool `json:"skip_holidays,omitempty"`

	// 考勤范围
	Offset *int `json:"offset,omitempty"`

	// 考勤地点相关设置信息
	Positions []AttendancePosition `json:"positions,omitempty"`

	// 默认班次ID
	DefaultClassId *int `json:"default_class_id,omitempty"`

	//是是否开启人脸检测。
	//
	//true：开启
	//
	//false：关闭（默认值）
	EnableFaceCheck *bool `json:"enable_face_check,omitempty"`

	// 蓝牙打卡相关配置信息
	BleDevices []BleDevice `json:"ble_device_list,omitempty"`

	//是否开启外勤打卡必须拍照。
	//
	//true：开启
	//
	//false：关闭（默认值）
	EnableOutsideCameraCheck *bool `json:"enable_outside_camera_check,omitempty"`

	// 自由工时考勤组工作日。
	FreeCheckWorkDays string `json:"freecheck_work_days,omitempty"`

	//是否可以外勤打卡。
	//
	//true：允许（默认值）
	//
	//false：不允许
	EnableOutsideCheck *bool `json:"enable_outside_check,omitempty"`

	// 周班次列表
	WorkdayClassList []string `json:"workday_class_list,omitempty"`

	// 考勤组名
	Name string `json:"name,omitempty" validate:"required"`

	//子管理员权限范围。
	//
	//w：可管理
	//
	//r：可读
	ResourcePermissions []ResourcePermission `json:"resource_permission_map,omitempty,omitempty"`

	// 企业id
	CorpId string `json:"corp_id,omitempty"`

	// 未排班时是否允许员工选择班次打卡
	EnableEmpSelectClass *bool `json:"enable_emp_select_class,omitempty"`
}

type BleDevice struct {
	// 设备ID，调用查询员工智能考勤机列表获取。
	DeviceId int `json:"device_id"`
}

// AttendancePosition 考勤地点相关设置信息
type AttendancePosition struct {
	// 考勤地址
	Address string `json:"address"`

	// 纬度
	Latitude string `json:"latitude,omitempty"`

	// 经度
	Longitude string `json:"longitude,omitempty"`

	// 精度
	Accuracy string `json:"accuracy,omitempty"`

	// 考勤标题
	Title string `json:"title,omitempty"`

	// 企业的corpId
	CorpId string `json:"corp_id,omitempty"`
}

type AttendanceShift struct {

	// 班次ID，可通过获取班次摘要信息接口获取。
	Id int `json:"id"`
}

type AttendanceWifi struct {

	// mac地址
	MacAddr string `json:"mac_addr"`

	// wifi的ssid。
	Ssid string `json:"ssid"`

	// 企业的corpId
	CorpId string `json:"corp_id"`
}

// AttendanceMember 考勤组成员相关设置信息
type AttendanceMember struct {
	// 角色，固定值Attendance。
	Role string `json:"role,omitempty" validate:"required"`

	// 用户userid。
	UserId string `json:"user_id,omitempty" validate:"required"`

	// 类型，固定值StaffMember
	Type string `json:"type,omitempty" validate:"required"`

	// 企业id
	CorpId string `json:"corp_id,omitempty,omitempty"`
}

func NewAttendanceMember(userId, corpId string) AttendanceMember {
	return AttendanceMember{"Attendance", userId, "StaffMember", corpId}
}

type ResourcePermission struct {

	// 员工排班
	Schedule string `json:"schedule,omitempty,omitempty" validate:"omitempty,oneof=w r"`

	// 设置外勤打卡
	OutSideCheck string `json:"out_side_check,omitempty" validate:"omitempty,oneof=w r"`

	// 设置考勤类型
	GroupType string `json:"group_type,omitempty" validate:"omitempty,oneof=w r"`

	// 设置参与考勤人员
	GroupMember string `json:"group_member,omitempty" validate:"omitempty,oneof=w r"`

	// 设置打卡方式
	CheckPositionType string `json:"check_position_type,omitempty" validate:"omitempty,oneof=w r"`

	// 设置拍照打卡规则
	CameraCheck string `json:"camera_check,omitempty" validate:"omitempty,oneof=w r"`

	// 设置考勤时间
	CheckTime string `json:"check_time,omitempty" validate:"omitempty,oneof=w r"`

	// 设置加班规则
	OverTimeRule string `json:"over_time_rule,omitempty" validate:"omitempty,oneof=w r"`
}

type CreateAttendanceGroupBuilder struct {
	group *CreateAttendanceGroup
}

func NewCreateAttendanceGroup(opUserId, name string, attendanceType attendance.Type, members []AttendanceMember) *CreateAttendanceGroupBuilder {
	return &CreateAttendanceGroupBuilder{group: &CreateAttendanceGroup{
		OpUserId:        opUserId,
		AttendanceGroup: &AttendanceGroup{Name: name, Type: string(attendanceType), Members: members},
	}}
}

func (b *CreateAttendanceGroupBuilder) SetOpUserId(userId string) *CreateAttendanceGroupBuilder {
	b.group.OpUserId = userId
	return b
}

func (b *CreateAttendanceGroupBuilder) SetCorpId(corpId string) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.CorpId = corpId
	return b
}

func (b *CreateAttendanceGroupBuilder) SetEnableEmpSelectClass(enable bool) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.EnableEmpSelectClass = &enable
	return b
}

func (b *CreateAttendanceGroupBuilder) SetEnableNextDay(enable bool) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.EnableNextDay = &enable
	return b
}

func (b *CreateAttendanceGroupBuilder) SetEnableOutsideCheck(enable bool) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.EnableOutsideCheck = &enable
	return b
}

func (b *CreateAttendanceGroupBuilder) SetDefaultClassId(id int) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.DefaultClassId = &id
	return b
}

func (b *CreateAttendanceGroupBuilder) SetModifyMember(modifyMember bool) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.ModifyMember = &modifyMember
	return b
}

func (b *CreateAttendanceGroupBuilder) SetCheckNeedHealthyCode(enable bool) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.CheckNeedHealthyCode = &enable
	return b
}

func (b *CreateAttendanceGroupBuilder) SetEnableFaceCheck(enable bool) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.EnableFaceCheck = &enable
	return b
}

func (b *CreateAttendanceGroupBuilder) SetDisableCheckWithoutSchedule(disable bool) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.DisableCheckWithoutSchedule = &disable
	return b
}

func (b *CreateAttendanceGroupBuilder) SetDisableCheckWhenRest(disable bool) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.DisableCheckWhenRest = &disable
	return b
}

func (b *CreateAttendanceGroupBuilder) SetAttendanceType(attendanceType attendance.Type) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.Type = string(attendanceType)
	return b
}

func (b *CreateAttendanceGroupBuilder) SetEnableCameraCheck(enable bool) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.EnableCameraCheck = &enable
	return b
}

func (b *CreateAttendanceGroupBuilder) SetEnablePositionBle(enable bool) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.EnablePositionBle = &enable
	return b
}

func (b *CreateAttendanceGroupBuilder) SetEnableOutsideCameraCheck(enable bool) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.EnableOutsideCameraCheck = &enable
	return b
}

func (b *CreateAttendanceGroupBuilder) SetName(name string) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.Name = name
	return b
}

func (b *CreateAttendanceGroupBuilder) SetBleDevices(deviceId int) *CreateAttendanceGroupBuilder {
	devices := b.group.AttendanceGroup.BleDevices
	b.group.AttendanceGroup.BleDevices = append(devices, BleDevice{deviceId})
	return b
}

func (b *CreateAttendanceGroupBuilder) SetWorkdayClassList(days []string) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.WorkdayClassList = days
	return b
}

func (b *CreateAttendanceGroupBuilder) SetResourcePermissions(schedule, outSideCheck, groupType, groupMember, checkPositionType, cameraCheck, checkTime, overTimeRule string) *CreateAttendanceGroupBuilder {
	permission := b.group.AttendanceGroup.ResourcePermissions
	b.group.AttendanceGroup.ResourcePermissions = append(permission,
		ResourcePermission{
			schedule,
			outSideCheck,
			groupType,
			groupMember,
			checkPositionType,
			cameraCheck,
			checkTime,
			overTimeRule,
		})
	return b
}

func (b *CreateAttendanceGroupBuilder) SetPositions(address, latitude, longitude, accuracy, title, corpId string) *CreateAttendanceGroupBuilder {
	positions := b.group.AttendanceGroup.Positions
	b.group.AttendanceGroup.Positions = append(positions,
		AttendancePosition{
			address,
			latitude,
			longitude,
			accuracy,
			title,
			corpId,
		})
	return b
}

func (b *CreateAttendanceGroupBuilder) SetOffset(offset int) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.Offset = &offset
	return b
}

func (b *CreateAttendanceGroupBuilder) SetFreeCheckDayStartMinOffset(start int) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.FreeCheckDayStartMinOffset = &start
	return b
}

func (b *CreateAttendanceGroupBuilder) SetSkipHolidays(skip bool) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.SkipHolidays = &skip
	return b
}

func (b *CreateAttendanceGroupBuilder) SetOwner(owner string) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.Owner = owner
	return b
}

func (b *CreateAttendanceGroupBuilder) SetAttendanceShifts(classId int) *CreateAttendanceGroupBuilder {
	shifts := b.group.AttendanceGroup.AttendanceShifts
	b.group.AttendanceGroup.AttendanceShifts = append(shifts, AttendanceShift{classId})
	return b
}

func (b *CreateAttendanceGroupBuilder) SetManagers(userId string, userIds ...string) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.Managers = append(userIds, userId)
	return b
}

func (b *CreateAttendanceGroupBuilder) SetSpecialDays(specialDays string) *CreateAttendanceGroupBuilder {
	b.group.AttendanceGroup.SpecialDays = specialDays
	return b
}

func (b *CreateAttendanceGroupBuilder) SetMembers(userId, corpId string) *CreateAttendanceGroupBuilder {
	members := b.group.AttendanceGroup.Members
	b.group.AttendanceGroup.Members = append(members,
		AttendanceMember{
			"Attendance",
			userId,
			"StaffMember",
			corpId,
		})
	return b
}

func (b *CreateAttendanceGroupBuilder) SetAttendanceWifi(addr, ssId, corpId string) *CreateAttendanceGroupBuilder {
	wifi := b.group.AttendanceGroup.Wifis
	b.group.AttendanceGroup.Wifis = append(wifi,
		AttendanceWifi{
			addr,
			ssId,
			corpId,
		})
	return b
}

func (b *CreateAttendanceGroupBuilder) Build() *CreateAttendanceGroup {
	return b.group
}

func (cg *CreateAttendanceGroup) String() string {
	if str, err := json.Marshal(cg); err != nil {
		return ""
	} else {
		return string(str)
	}
}
