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

package model

//文件消息
type file struct {
	MediaId string `json:"media_id" validate:"required"` //媒体文件ID。引用的媒体文件最大10MB。可以通过上传媒体文件接口获取。
}

type fileMessage struct {
	message
	file `json:"file" validate:"required"`
}

func NewFileMessage(mediaId string) fileMessage {
	return fileMessage{message: message{MsgType: "file"}, file: file{MediaId: mediaId}}
}

////请求参数验证
//func (req fileMessage) Validate(valid *validator.Validate, trans translator.Translator) error {
//	if err := valid.Struct(req); err != nil {
//		errs := err.(validator.ValidationErrors)
//		var slice []string
//		for _, msg := range errs {
//			slice = append(slice, msg.Translate(trans))
//		}
//		return errors.New(strings.Join(slice, ","))
//	}
//	return nil
//}
