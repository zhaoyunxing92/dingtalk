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

type RegisterEvent struct {
	Response
	// 注册的事件类型
	Tags []string `json:"call_back_tag"`

	// 加解密需要用到的token，可以随机填写，长度大于等于6个字符且少于64个字符。
	Token string `json:"token"`

	// 接收事件回调的url，必须是公网可以访问的url地址。
	Url string `json:"url"`

	//数据加密密钥。用于回调数据的加密，长度固定为43个字符，从a-z，A-Z，0-9共62个字符中选取，您可以随机生成，ISV(服务提供商)推荐使用注册套件时填写的EncodingAESKey
	Secret string `json:"aes_key"`
}
