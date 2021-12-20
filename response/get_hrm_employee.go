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

type GetHrmEmployee struct {
	Response

	getHrmEmployee `json:"result"`
}

type getHrmEmployee struct {

	// 查询到的员工userid
	UserIds []string `json:"data_list"`

	// 下一次分页调用的offset值，当返回结果里没有next_cursor时，表示分页结束
	NextCursor int `json:"next_cursor"`
}
