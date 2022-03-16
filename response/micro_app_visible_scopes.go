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

// MicroAppVisibleScopes 应用可见范围
type MicroAppVisibleScopes struct {
	Response

	// 应用可见的用户列表。
	UserVisibleScopes []string `json:"userVisibleScopes"`

	// 应用可见的部门列表。
	DeptVisibleScopes []int `json:"deptVisibleScopes"`

	// 是否仅限管理员可见，true代表仅限管理员可见。
	Hidden bool `json:"isHidden"`
}
