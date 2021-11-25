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

package dingtalk

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/domain/message"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
)

func TestDingTalk_CreateChat(t *testing.T) {
	t.Skip()
	res, err := client.CreateChat(
		request.NewCreatChat("dingtalk", "manager164",
			"manager164", "manager164", "manager164").
			Build())

	// chat79e713edce46ea72ea0dd60fb4f9d6f6
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetChatInfo(t *testing.T) {
	res, err := client.GetChatInfo("chat8ff884ef696f5717678c6280edfdbbf1")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.Name, "ci测试群")
}

func TestDingTalk_UpdateChat(t *testing.T) {

	res, err := client.UpdateChat(
		request.NewUpdateChat("chat8ff884ef696f5717678c6280edfdbbf1").
			SetName("ci测试群").
			SetShowHistory(0).
			SetSearchable(1).
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetChatQRCode(t *testing.T) {

	res, err := client.GetChatQRCode("chat8ff884ef696f5717678c6280edfdbbf1", "manager164")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res.Url)
}

func TestDingTalk_ChatFriendSwitch(t *testing.T) {

	res, err := client.ChatFriendSwitch("chat8ff884ef696f5717678c6280edfdbbf1", true)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_ChatSetUserNick(t *testing.T) {

	res, err := client.ChatSetUserNick("chat8ff884ef696f5717678c6280edfdbbf1", "manager164", "小二")

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_SendMsgToChat(t *testing.T) {
	msg := message.NewOaMessage("消息的头部标题", "FFBBBBBB", "http://dingtalk.com",
		"http://dingtalk.com")

	msg.Body.Title = "消息体标题消息体标题"
	msg.Body.Content = "消息体的内容，最多显示3行"
	msg.Body.MediaId = "@lADOADmaWMzazQKA"
	msg.Body.Forms = []message.Form{
		{Key: "姓名", Value: "张三"},
		{Key: "年龄", Value: "20"},
		{Key: "身高", Value: "1.8m"},
		{Key: "体重", Value: "70kg"},
		{Key: "学历", Value: "本科"},
		{Key: "爱好", Value: "打球、听歌"},
	}
	msg.Rich.Num = "15.6"
	msg.Rich.Unit = "元"

	msg.Author = "李四"

	msg.StatusBar.Value = "进行中"
	msg.StatusBar.BackColor = "0xFFF65E5E"

	res, err := client.SendChatMessage("chat8ff884ef696f5717678c6280edfdbbf1", msg)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
