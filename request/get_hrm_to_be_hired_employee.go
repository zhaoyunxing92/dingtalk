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

type GetHrmToBeHiredEmployee struct {
	// 分页游标，从0开始。根据返回结果里的next_cursor是否为空来判断是否还有下一页，且再次调用时offset设置成next_cursor的值。
	Offset int `json:"offset"`

	// 分页大小，最大50。
	Size int `json:"size" validate:"max=50"`
}

func NewGetHrmToBeHiredEmployee(offset, size int) *GetHrmToBeHiredEmployee {
	return &GetHrmToBeHiredEmployee{offset, size}
}
