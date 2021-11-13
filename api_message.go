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

package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/domain"
	"net/http"
)

//SendToConversation:发送普通消息
func (ding *dingTalk) SendToConversation(senderId, chatId string, msg domain.Request) (req domain.SendToConversationResponse, err error) {

	form := make(map[string]interface{}, 3)
	form["sender"] = senderId
	form["cid"] = chatId
	form["msg"] = msg

	err = ding.Request(http.MethodPost, constant.SendToConversationKey, nil, form, &req)
	return req, err
}
