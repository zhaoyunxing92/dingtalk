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

import "testing"
import "github.com/zhaoyunxing92/dingtalk/v2/request"
import "github.com/stretchr/testify/assert"

func TestDingTalk_SendTemplateMessage(t *testing.T) {

	res, err := isv.SendTemplateMessage(
		request.NewSendTemplateMessage(1332307896, "80424a44cb444c53a071aae34c0fd140").
			SetUserIds("manager7556", "manager7556").
			SetData("msg", "发送模板消息测试").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetMessageProgress(t *testing.T) {

	res, err := isv.GetMessageProgress(1332307896, 474203371488)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_SendCorpConversationMessage_Text(t *testing.T) {

	res, err := isv.SendCorpConversationMessage(request.NewTextMessage("hello"))

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
