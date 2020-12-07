package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/global"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
)

//Translate：文本翻译
func (talk *DingTalk) Translate(query, sourceLanguage, targetLanguage string) (rsp model.DeptCreateResponse, err error) {

	form := make(map[string]interface{}, 2)
	form["query"] = query
	form["source_language"] = sourceLanguage
	form["target_language"] = targetLanguage

	err = talk.request(http.MethodPost, global.CreateDeptKey, nil, form, &rsp)

	return rsp, err
}
