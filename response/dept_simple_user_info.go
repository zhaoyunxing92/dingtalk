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

type DeptSimpleUserInfo struct {
	Response
	deptUserInfoPage `json:"result"`
}

type deptUserInfoPage struct {
	//是否还有更多的数据
	HasMore bool `json:"has_more"`

	//下一次分页的游标，如果has_more为false，表示没有更多的分页数据。
	NextCursor int `json:"next_cursor"`

	DeptUsers []simpleUserInfo `json:"list"`
}

//simpleUserInfo 简单用户信息
type simpleUserInfo struct {
	UserId string `json:"userid"`

	Name string `json:"name"`
}
