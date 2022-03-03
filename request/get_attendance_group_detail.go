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

type GetAttendanceGroupDetail struct {

	// 操作人userid
	UserId string `json:"op_user_id"`

	//考勤组ID。
	//
	//说明 如果你使用的是旧考勤组标识即group_key，
	//企业内部应用可以调用groupKey转换为groupId接口将group_key转换为group_id。
	//钉钉三方企业应用可以调用groupKey转换为groupId接口将group_key转换为group_id
	GroupId int `json:"group_id"`
}

func NewGetAttendanceGroupDetail(userId string, groupId int) *GetAttendanceGroupDetail {
	return &GetAttendanceGroupDetail{userId, groupId}
}
