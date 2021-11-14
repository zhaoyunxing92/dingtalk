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

import "io"

type MediaUpload struct {
	//媒体文件类型：
	//
	//image：图片，图片最大1MB。支持上传jpg、gif、png、bmp格式。
	//
	//voice：语音，语音文件最大2MB。支持上传amr、mp3、wav格式。
	//
	//video：视频，视频最大10MB。支持上传mp4格式。
	//
	//file：普通文件，最大10MB。支持上传doc、docx、xls、xlsx、ppt、pptx、zip、pdf、rar格式。
	Type string `json:"type" validate:"required,oneof=image voice file video"`

	//要上传的媒体文件
	Media fileItem `json:"media" validate:"required"`

	// 要上传的文件路径
	Path string `json:"-"  validate:"required"`

	Reader io.Reader `validate:"required"`
}

type fileItem struct {
	//文件名称
	FileName string `validate:"required"`

	//字段名称
	FieldName string `validate:"required"`
}
