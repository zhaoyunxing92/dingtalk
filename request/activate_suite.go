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

type ActivateSuite struct {
	//第三方应用的SuiteKey
	SuiteKey string `json:"suite_key,omitempty" validate:"required"`

	//授权企业的CorpId。
	CorpId string `json:"auth_corpid,omitempty" validate:"required"`

	//授权企业的永久授权码
	PermanentCode string `json:"permanent_code,omitempty" validate:"required"`
}

func NewActivateSuite(suiteKey, corpId, permanentCode string) *ActivateSuite {
	return &ActivateSuite{suiteKey, corpId, permanentCode}
}
