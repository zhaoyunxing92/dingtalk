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

type GetAttendanceLeaveStatus struct {
	// 待查询用户的ID列表，每次最多100个
	UserIdList string `json:"userid_list"`

	// 开始时间 ，Unix时间戳，支持最多180天的查询
	StartTime int64 `json:"start_time"`

	// 结束时间，Unix时间戳，支持最多180天的查询
	EndTime int64 `json:"end_time"`

	// 支持分页查询，与size参数同时设置时才生效，此参数代表偏移量，偏移量从0开始
	Offset int64 `json:"offset"`

	// 支持分页查询，与offset参数同时设置时才生效，此参数代表分页大小，最大20
	Size int64 `json:"size"`
}
