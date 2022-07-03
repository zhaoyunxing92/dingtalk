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

type RoleUser struct {
	Response

	Result struct {
		HasMore    bool `json:"hasMore"`
		NextCursor int  `json:"nextCursor"`

		Users []struct {
			UserId string `json:"userid"`

			Name string `json:"name"`

			ManageScopes []struct {
				// 部门Id
				DeptId int `json:"dept_id"`
				// 部门名称
				DeptName string `json:"name"`
			} `json:"manageScopes"`
		} `json:"list"`
	} `json:"result"`
}
