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

type GetAttendanceLeaveStatus struct {
	Response

	AttendanceLeaveStatus struct {
		HasMore     bool            `json:"has_more"`     // 是否有更多数据。true：有，false：没有
		LeaveStatus []LeaveStatusVO `json:"leave_status"` // 请假状态列表。
	} `json:"result"`
}

// LeaveStatusVO 请假状态
type LeaveStatusVO struct {
	DurationUnit    string `json:"duration_unit"`    // 请假单位：percent_day：天，percent_hour：小时
	DurationPercent int64  `json:"duration_percent"` // 假期时长*100，例如用户请假时长为1天，该值就等于100
	EndTime         int64  `json:"end_time"`         // 请假结束时间，Unix时间戳
	StartTime       int64  `json:"start_time"`       // 请假开始时间，Unix时间戳
	UserId          string `json:"userid"`           // 用户ID
	LeaveCode       string `json:"leave_code"`       // 这个文档里面没有，但是实际接口返回有, 应该是请假类型
}
