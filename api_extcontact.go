package dingtalk

import (
	"net/http"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

//CreateExtContact 添加外部联系人
func (ding *dingTalk) CreateExtContact(res *request.CreateExtContact) (rsp response.CreateExtContact, err error) {

	return rsp, ding.Request(http.MethodPost, constant.CreateExtContactKey, nil, res, &rsp)
}

//DeleteExtContact 删除外部联系人
func (ding *dingTalk) DeleteExtContact(userId string) (rsp response.Response, err error) {

	return rsp, ding.Request(http.MethodPost, constant.DeleteExtContactKey, nil,
		request.NewDeleteExtContact(userId), &rsp)
}

//UpdateExtContact 更新外部联系人
func (ding *dingTalk) UpdateExtContact(res *request.UpdateExtContact) (rsp response.CreateExtContact, err error) {

	return rsp, ding.Request(http.MethodPost, constant.UpdateExtContactKey, nil, res, &rsp)
}

//GetExtContact 获取外部联系人列表
func (ding *dingTalk) GetExtContact(offset, size int) (rsp response.GetExtContact, err error) {

	return rsp, ding.Request(http.MethodPost, constant.GetExtContactKey, nil,
		request.NewGetExtContact(offset, size), &rsp)
}

//GetExtContactLabel 获取外部联系人标签列表
func (ding *dingTalk) GetExtContactLabel(offset, size int) (rsp response.GetExtContactLabel, err error) {

	return rsp, ding.Request(http.MethodPost, constant.GetExtContactLabelKey, nil,
		request.NewGetExtContactLabel(offset, size), &rsp)
}

//GetExtContactDetail 获取外部联系人详情
func (ding *dingTalk) GetExtContactDetail(userId string) (rsp response.GetExtContactDetail, err error) {

	return rsp, ding.Request(http.MethodPost, constant.GetExtContactDetailKey, nil,
		request.NewGetExtContactDetail(userId), &rsp)
}
