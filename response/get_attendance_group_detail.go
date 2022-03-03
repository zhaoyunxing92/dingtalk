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

type GetAttendanceGroupDetail struct {
	Response

	AttendanceGroup struct {
		// 考勤地址
		AddressList []string `json:"address_list"`

		// 考勤组ID
		Id int `json:"id"`

		// 人员人数
		MemberCount int `json:"member_count"`

		// 考勤组名称
		Name string `json:"name"`

		// 考勤组主负责人的userid
		OwnerUserId string `json:"owner_user_id"`

		// 排班ID
		ShiftIds []int `json:"shift_ids"`

		//考勤组类型：
		//
		//FIXED：代表固定班制考勤组
		//
		//TURN：代表排班制考勤组
		//
		//NONE：代表自由工时考勤组
		Type string `json:"type"`

		// 跳转链接
		Url string `json:"url"`

		// wifi名称
		Wifis []string `json:"wifis"`

		// 工作日
		WorkDayList []int `json:"work_day_list"`
	} `json:"result"`
}
