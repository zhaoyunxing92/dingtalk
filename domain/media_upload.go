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

package domain

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"io"
	"strings"
)

//媒体文件类型：
//image：图片，图片最大1MB，支持JPG格式。
//voice：语音，语音文件最大2MB，支持AMR格式。
//file：普通文件，最大10MB。支持上传以下文件：doc、docx、xls、xlsx、ppt、pptx、zip、pdf、rar。
type MediaUpload struct {
	Response
	Type      string `json:"type"`       // 媒体文件类型
	MediaId   string `json:"media_id"`   // 媒体文件类型
	CreatedAt int    `json:"created_at"` // 媒体文件类型
}

//请求
type UploadFile struct {
	Type      string    `validate:"required,oneof=image voice file"` // 媒体文件类型
	FileName  string    `validate:"required"`                        // 文件名称
	FieldName string    `validate:"required"`                        //字段名称
	Reader    io.Reader `validate:"required"`
}

func NewUploadFile(name, tp string, reader io.Reader) UploadFile {
	return UploadFile{tp, name, "media", reader}
}

//请求参数验证
func (req *UploadFile) Validate(valid *validator.Validate, trans translator.Translator) error {
	if err := valid.Struct(req); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Translate(trans))
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}
