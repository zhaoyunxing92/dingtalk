package dingtalk

import (
	"net/http"
	"net/url"

	"github.com/lihongchen/dingtalk/global"
	"github.com/lihongchen/dingtalk/model"
)

func (talk *DingTalk) MediaUpload(req model.UploadFile) (media model.MediaUpload, err error) {

	//参数验证
	if err = req.Validate(talk.validate, talk.trans); err != nil {
		return model.MediaUpload{}, err
	}
	params := url.Values{}
	params.Add("type", req.Type)

	err = talk.request(http.MethodPost, global.MediaUploadKey, params, req, &media)
	return media, err
}
