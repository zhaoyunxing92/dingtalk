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

type UnactiveCorp struct {
	Response

	// 微应用ID
	AppId int `json:"app_id"`

	// 授权未激活应用的CorpId列表
	CorpIds []string `json:"corp_list"`

	// 是否还有更多数据
	HasMore bool `json:"has_more"`
}
