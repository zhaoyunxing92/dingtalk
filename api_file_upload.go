package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/model"
	"net/http"
	"net/url"
)

func (ding *DingTalk) MediaUpload(req model.UploadFile) (media model.MediaUpload, err error) {

	params := url.Values{}
	params.Add("type", req.Type)

	err = ding.Request(http.MethodPost, constant.MediaUploadKey, params, req, &media)
	return media, err
}
