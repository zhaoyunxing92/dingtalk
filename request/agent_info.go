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

//AgentInfo 获取授权应用的基本信息
type AgentInfo struct {
	AgentId  int    `json:"agentid" validate:"required"`
	SuiteKey string `json:"suite_key" validate:"required"`
	CorpId   string `json:"auth_corpid" validate:"required"`
}

func NewAgentInfo(agentId int, key, corpId string) *AgentInfo {
	return &AgentInfo{agentId, key, corpId}
}
