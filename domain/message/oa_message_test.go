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

package message

import (
	"testing"
)

func TestNewOaMessage(t *testing.T) {

	msg := NewOaMessage("消息的头部标题", "FFBBBBBB", "http://dingtalk.com",
		"http://dingtalk.com")

	msg.Body.Title = "消息体"
	msg.Body.Content = "消息体的内容，最多显示3行"
	msg.Body.MediaId = "@lADOADmaWMzazQKA"
	msg.Body.Forms = []Form{
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

	t.Log(msg.String())
}
