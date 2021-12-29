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

type GetAttendanceGroup struct {
	// 支持分页查询，与size参数同时设置时才生效，此参数代表偏移量，偏移量从0开始，下次调用传上次调用时的size与offset之和
	Offset int `json:"offset"`

	// 分页大小，最大10。
	Size int `json:"size" validate:"max=10"`
}

func NewGetAttendanceGroup(offset, size int) *GetAttendanceGroup {
	return &GetAttendanceGroup{offset, size}
}
