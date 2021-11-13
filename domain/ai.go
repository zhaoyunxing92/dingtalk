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

//翻译返回
type AiResponse struct {
	Response
	Result string `json:"result"` //翻译结果字符串
}

//OcrStructuredResult
type OcrStructuredResult struct {
	Height         int    `json:"height"`          //旋转后图片高度。
	Width          int    `json:"width"`           //旋转后图片宽度。
	Angle          int    `json:"angle"`           //旋转度。
	Data           string `json:"data"`            //图片识别内容json字符串。
	OriginalHeight int    `json:"original_height"` //图片识别内容json字符串。
	OriginalWidth  int    `json:"original_width"`  //图片识别内容json字符串。
}

type OcrStructuredResponse struct {
	Response
	Result OcrStructuredResult `json:"result"` //翻译结果字符串
}
