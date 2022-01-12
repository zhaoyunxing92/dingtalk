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

type FetchDeptList struct {
	Response

	// 部门
	Dept []struct {

		//当部门群已经创建后，是否有新人加入部门会自动加入该群：
		//
		//true：自动加入群
		//
		//false：不会自动加入群
		AutoAddUser bool `json:"autoAddUser"`

		//是否同步创建一个关联此部门的企业群：
		//
		//true：创建
		//
		//false：不创建
		CreateDeptGroup bool `json:"createDeptGroup"`

		// 部门id
		Id int `json:"id"`

		// 部门名称
		Name string `json:"name"`

		// 父部门id
		ParentId int `json:"parentid"`
	} `json:"department"`
}
