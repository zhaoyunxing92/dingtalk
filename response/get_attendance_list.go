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

type GetAttendanceList struct {
	Response

	AttendanceList []struct {
		// 数据来源：
		// - ATM：考勤机打卡（指纹/人脸打卡）
		// - BEACON：IBeacon
		// - DING_ATM：钉钉考勤机（考勤机蓝牙打卡）
		// - USER：用户打卡
		// - BOSS：老板改签
		// - APPROVE：审批系统
		// - SYSTEM：考勤系统
		// - AUTO_CHECK：自动打卡
		SourceType string `json:"sourceType"`

		// 计算迟到和早退，基准时间
		BaseCheckTime int64 `json:"baseCheckTime"`

		// 实际打卡时间, 用户打卡时间的毫秒数
		UserCheckTime int64 `json:"userCheckTime"`

		// 关联的审批实例ID，当该字段非空时，表示打卡记录与请假、加班等审批有关
		ProcInstId string `json:"procInstId"`

		// 关联的审批ID，当该字段非空时，表示打卡记录与请假、加班等审批有关
		ApproveId int64 `json:"approveId"`

		// 位置结果：
		// - Normal：范围内
		// - Outside：范围外
		// - NotSigned：未打卡
		LocationResult string `json:"locationResult"`

		// 打卡结果：
		// - Normal：正常
		// - Early：早退
		// - Late：迟到
		// - SeriousLate：严重迟到
		// - Absenteeism：旷工迟到
		// - NotSigned：未打卡
		TimeResult string `json:"timeResult"`

		// 考勤类型：
		// - OnDuty：上班
		// - OffDuty：下班
		CheckType string `json:"checkType"`

		// 打卡人的userId
		UserId string `json:"userId"`

		// 工作日
		WorkDate int64 `json:"workDate"`

		// 打卡记录ID
		RecordId int64 `json:"recordId"`

		// 排班ID
		PlanId int64 `json:"planId"`

		// 考勤组ID
		GroupId int64 `json:"groupId"`

		// 唯一标识ID
		Id int64 `json:"id"`
	} `json:"recordresult"`
}
