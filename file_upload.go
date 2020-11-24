package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/domain"
	"github.com/zhaoyunxing92/dingtalk/global"
	"net/http"
	"net/url"
)

func (talk *DingTalk) MediaUpload(req domain.UploadFile) (media domain.MediaUpload, err error) {

	//参数验证
	if err = req.Validate(talk.validate, talk.trans); err != nil {
		return domain.MediaUpload{}, err
	}
	params := url.Values{}
	params.Add("type", req.Type)

	err = talk.request(http.MethodPost, global.MediaUploadKey, params, req, &media)
	return media, err
}
