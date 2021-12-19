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

type GetIndustryDept struct {
	DeptId int `json:"dept_id" validate:"required"`

	// 支持分页查询，与size参数同时设置时才生效，此参数代表偏移量，偏移量从0开始
	Offset int `json:"cursor"`

	// 支持分页查询，与offset参数同时设置时才生效，此参数代表分页大小，最大100
	Size int `json:"size" validate:"max=1000"`
}

func NewIndustryDept(deptId, offset, size int) *GetIndustryDept {
	return &GetIndustryDept{deptId, offset, size}
}
