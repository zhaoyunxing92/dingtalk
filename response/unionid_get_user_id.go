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

type UnionIdGetUserId struct {
	Response

	UserId `json:"result"`
}

type UserId struct {
	UserId string `json:"userid"`

	//联系类型：
	//
	//0：企业内部员工
	//
	//1：企业外部联系人
	ContactType int `json:"contact_type"`
}
