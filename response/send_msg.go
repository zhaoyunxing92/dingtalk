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

type SendMessage struct {
	Response

	//有效接收消息的员工的userid。
	//
	//接收者可以是单聊接收者或者群聊会话里的接收者，如果接收者是当前接口调用所使用的企业的员工，则是有效接收者。
	//
	//接口返回所有有效接收者的userid。
	//
	//非有效接收者是收不到消息的。
	Receiver string `json:"receiver"`
}
