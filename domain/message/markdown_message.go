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

package message

type markdown struct {
	//首屏会话透出的展示内容
	Title string `json:"title" validate:"required"`

	//markdown格式的消息，建议500字符以内
	Text string `json:"text" validate:"required"`
}

//markdown消息
type markdownMessage struct {
	message
	markdown `json:"markdown" validate:"required"`
}

func (v *markdownMessage) MessageType() string {
	return "markdown"
}

func NewMarkDownMessage(title, content string) *markdownMessage {
	return &markdownMessage{message{MsgType: "markdown"}, markdown{title, content}}
}
