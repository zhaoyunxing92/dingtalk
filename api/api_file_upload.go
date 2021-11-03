package api

import (
	"github.com/zhaoyunxing92/dingtalk/constant"
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
	"net/url"
)

func (ding *DingTalk) MediaUpload(req model.UploadFile) (media model.MediaUpload, err error) {

	params := url.Values{}
	params.Add("type", req.Type)

	err = ding.request(http.MethodPost, constant.MediaUploadKey, params, req, &media)
	return media, err
}
