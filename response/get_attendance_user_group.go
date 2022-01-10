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

type GetAttendanceUserGroup struct {
	Response

	// 查询结果
	Result struct {

		// 考勤组ID
		GroupId int `json:"group_id"`

		// 考勤组名称
		Name string `json:"name"`

		//考勤类型：
		//
		//FIXED：固定排班
		//
		//TURN：轮班排班
		//
		//NONE：无班次
		Type string `json:"type"`

		// 考勤组中的班次列表
		Classes []struct {
			// 班次ID
			ClassId int `json:"class_id"`

			// 班次名称
			Name string `json:"name"`

			// 班次中上下班列表
			Sections []struct {
				// 班次中上下班详情列表
				Times []struct {
					// 打卡跨越的时间天数
					Across int `json:"across"`

					// 打卡时间
					CheckTime string `json:"check_time"`

					//打卡类型：
					//
					//OnDuty：上班
					//
					//OffDuty：下班
					CheckType string `json:"check_type"`

					// 允许的最晚延后打卡时间，单位分钟
					EndMin int `json:"end_min"`

					// 允许的最早提前打卡时间，单位分钟
					BeginMin int `json:"begin_min"`
				} `json:"times"`
			} `json:"sections"`

			// 班次配置
			Setting struct {

				// 休息开始设置
				RestBeginTime struct {

					//时间跨度。
					//
					//0：不跨天
					//
					//1：跨天
					Across int `json:"across"`

					// 休息开始时间
					CheckTime string `json:"check_time"`

					//设置类型：
					//
					//OnDuty：休息开始
					//
					//OffDuty：休息结束
					CheckType string `json:"check_type"`
				} `json:"rest_begin_time"`

				// 休息结束时间设置
				RestEndTime struct {

					// 时间跨度
					Across int `json:"across"`

					// 休息结束时间
					CheckTime string `json:"check_time"`

					//打卡类型：
					//
					//OnDuty：休息开始
					//
					//OffDuty：休息结束
					CheckType string `json:"check_type"`
				} `json:"rest_end_time"`
			} `json:"setting"`
		} `json:"classes"`
	} `json:"result"`
}
