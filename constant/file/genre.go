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

package file

type Genre string

const (
	// Image 图片，图片最大1MB。支持上传jpg、gif、png、bmp格式
	Image = Genre("image")

	// Voice 语音，语音文件最大2MB。支持上传amr、mp3、wav格式
	Voice = Genre("voice")

	// Video 视频，视频最大10MB。支持上传mp4格式
	Video = Genre("video")

	// File 普通文件，最大10MB。支持上传doc、docx、xls、xlsx、ppt、pptx、zip、pdf、rar格式
	File = Genre("file")

	// Folder 文件夹
	Folder = Genre("folder")
)
