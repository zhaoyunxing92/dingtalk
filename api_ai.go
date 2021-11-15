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
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/domain"
	"net/http"
)

//Translate：文本翻译
func (ding *dingTalk) Translate(query, sourceLanguage, targetLanguage string) (rsp domain.AiResponse, err error) {

	form := make(map[string]string, 3)
	form["query"] = query
	form["source_language"] = sourceLanguage
	form["target_language"] = targetLanguage

	err = ding.Request(http.MethodPost, constant.TranslateKey, nil, form, &rsp)

	return rsp, err
}

//OcrRecognize：OCR文字识别
//ocrType:识别类型
//imageUrl:图片url
func (ding *dingTalk) OcrRecognize(ocrType, imageUrl string) (rsp domain.OcrStructuredResponse, err error) {

	//if ocrType != "idcard" || ocrType != "invoice" || ocrType != "blicense" || ocrType != "bank_card" ||
	//	ocrType != "car_no" || ocrType != "car_invoice" || ocrType != "driving_license" || ocrType != "vehicle_license" ||
	//	ocrType != "train_ticket" || ocrType != "quota_invoice" || ocrType != "taxi_ticket" || ocrType != "air_itinerary" ||
	//	ocrType != "approval_table" || ocrType != "roster" {
	//	err = errors.New(fmt.Sprintf("识别图片类型:[%s]不支持", ocrType))
	//}

	form := make(map[string]string, 2)
	form["type"] = ocrType
	form["image_url"] = imageUrl

	err = ding.Request(http.MethodPost, constant.OrcRecognizeKey, nil, form, &rsp)
	return rsp, err
}

//VoiceTranslate：ASR 一句话语音识别
//mediaId:音频id
func (ding *dingTalk) VoiceTranslate(mediaId string) (rsp domain.AiResponse, err error) {

	form := make(map[string]string, 1)
	form["media_id"] = mediaId

	err = ding.Request(http.MethodPost, constant.VoiceTranslateKey, nil, form, &rsp)
	return rsp, err
}
