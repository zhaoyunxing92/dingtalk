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

type GetAttendanceGetUpdateData struct {
	Response

	AttendanceGetUpdateData struct {
		WorkDate             string                        `json:"work_date"`              // 查询日期
		AttendanceResultList []AtAttendanceResultForOpenVo `json:"attendance_result_list"` // 打卡结果
		UserId               string                        `json:"userid"`                 // 用户userId
		ApproveList          []AtApproveForOpenVo          `json:"approve_list"`           // 审批单列表
		CheckRecordList      []AtAttendanceRecordForOpenVo `json:"check_record_list"`      // 打卡详情
		CorpId               string                        `json:"corpId"`                 // 企业corpId
		ClassSettingInfo     *AtClassSettingInfoForOpenVo  `json:"class_setting_info"`     // 当前排班对应的休息时间段
	} `json:"result"`
}

// 打卡结果
type AtAttendanceResultForOpenVo struct {
	RecordId       int64  `json:"record_id"`       // 打卡流水ID
	SourceType     string `json:"source_type"`     // 打卡来源
	PlanCheckTime  string `json:"plan_check_time"` // 标准打卡时间
	ClassId        int64  `json:"class_id"`        // 班次ID
	LocationMethod string `json:"location_method"` // 定位方法
	LocationResult string `json:"location_result"` // 定位结果
	OutsideRemark  string `json:"outside_remark"`  // 外勤备注
	PlanId         int64  `json:"plan_id"`         // 排班ID
	UserAddress    string `json:"user_address"`    // 用户打卡地址
	GroupId        int64  `json:"group_id"`        // 考勤组ID
	UserCheckTime  string `json:"user_check_time"` // 用户打卡时间
	ProcInstId     string `json:"procInst_id"`     // 审批单ID
	CheckType      string `json:"check_type"`      // 打卡类型
	TimeResult     string `json:"time_result"`     // 打卡的时间结果
	UserId         string `json:"userid"`          // 用户userId
}

// 审批单
type AtApproveForOpenVo struct {
	DurationUnit string `json:"duration_unit"` // 审批单的单位
	Duration     string `json:"duration"`      // 时长
	SubType      string `json:"sub_type"`      // 子类型名称
	TagName      string `json:"tag_name"`      // 审批单类型名称支持类型如下：请假、出差、外出、加班
	ProcInstId   string `json:"procInst_id"`   // 审批单ID
	BeginTime    string `json:"begin_time"`    // 审批单开始时间
	BizType      int64  `json:"biz_type"`      // 审批单类型：1-加班，2-出差/外出，3-请假
	EndTime      string `json:"end_time"`      // 审批单结束时间
	GmtFinished  string `json:"gmt_finished"`  // 审批单审批完成时间
}

// 打卡详情
type AtAttendanceRecordForOpenVo struct {
	RecordId          int64  `json:"record_id"`           // 打卡记录ID
	SourceType        string `json:"source_type"`         // 打卡来源
	UserAccuracy      string `json:"user_accuracy"`       // 用户定位精度
	ValidMatched      bool   `json:"valid_matched"`       // 是否匹配打卡结果的流水
	UserCheckTime     string `json:"user_check_time"`     // 用户打卡时间
	UserLongitude     string `json:"user_longitude"`      // 打卡经度
	UserSsid          string `json:"user_ssid"`           // wifi名称
	BaseAccuracy      string `json:"base_accuracy"`       // 基本定位精度
	UserMacAddr       string `json:"user_mac_addr"`       // MAC地址
	UserLatitude      string `json:"user_latitude"`       // 打卡纬度
	BaseAddress       string `json:"base_address"`        // 打卡基础地址
	InvalidRecordMsg  string `json:"invalid_record_msg"`  // 打卡无效的原因
	InvalidRecordType string `json:"invalid_record_type"` // 打卡无效的类型
}

// 班次休息时间段
type AtClassSettingInfoForOpenVo struct {
	RestTimeVoList []AtRestTimeVo `json:"rest_time_vo_list"` // 班次内休息信息
}

// 休息时间
type AtRestTimeVo struct {
	RestEndTime   int64 `json:"rest_end_time"`   // 休息结束时间
	RestBeginTime int64 `json:"rest_begin_time"` // 休息开始时间
}
