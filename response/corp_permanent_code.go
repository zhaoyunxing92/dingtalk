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

type CorpPermanentCode struct {
	Response
	corpAuthInfo  `json:"auth_corp_info"`
	PermanentCode string `json:"permanent_code"`
}

type corpAuthInfo struct {
	// 授权方企业CorpId
	CorpId string `json:"corpid"`

	// 授权方企业名称。
	CorpName string `json:"corp_name"`
}
