/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package request

type ReauthCorp struct {

	//套件下的微应用ID
	AppId int `json:"app_id" validate:"required,min=1"`

	//授权未激活应用的CorpId列表
	CorpIds []string `json:"corpid_list" validate:"required,min=1"`
}

func NewReauthCorp(appId int, corpIds []string) *ReauthCorp {
	return &ReauthCorp{appId, removeStringDuplicates(corpIds)}
}
