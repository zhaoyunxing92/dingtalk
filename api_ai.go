package dingtalk

import (
	"errors"
	"fmt"
	"github.com/zhaoyunxing92/dingtalk/global"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
)

//Translate：文本翻译
func (talk *DingTalk) Translate(query, sourceLanguage, targetLanguage string) (rsp model.AiResponse, err error) {

	form := make(map[string]string, 3)
	form["query"] = query
	form["source_language"] = sourceLanguage
	form["target_language"] = targetLanguage

	err = talk.request(http.MethodPost, global.TranslateKey, nil, form, &rsp)

	return rsp, err
}

//OcrRecognize：OCR文字识别
//ocrType:识别类型
//imageUrl:图片url
func (talk *DingTalk) OcrRecognize(ocrType, imageUrl string) (rsp model.OcrStructuredResponse, err error) {

	if ocrType != "idcard" || ocrType != "invoice" || ocrType != "blicense" || ocrType != "bank_card" ||
		ocrType != "car_no" || ocrType != "car_invoice" || ocrType != "driving_license" || ocrType != "vehicle_license" ||
		ocrType != "train_ticket" || ocrType != "quota_invoice" || ocrType != "taxi_ticket" || ocrType != "air_itinerary" ||
		ocrType != "approval_table" || ocrType != "roster" {
		err = errors.New(fmt.Sprintf("识别图片类型:[%s]不支持", ocrType))
	}

	form := make(map[string]string, 2)
	form["type"] = ocrType
	form["image_url"] = imageUrl

	err = talk.request(http.MethodPost, global.OrcRecognizeKey, nil, form, &rsp)
	return rsp, err
}

//VoiceTranslate：ASR 一句话语音识别
//mediaId:音频id
func (talk *DingTalk) VoiceTranslate(mediaId string) (rsp model.AiResponse, err error) {

	form := make(map[string]string, 1)
	form["media_id"] = mediaId

	err = talk.request(http.MethodPost, global.VoiceTranslateKey, nil, form, &rsp)
	return rsp, err
}
