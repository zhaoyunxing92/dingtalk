package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/crypto"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// GetAuthInfo 获取企业授权信息
func (ding *dingTalk) GetAuthInfo(corpId string) (string, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	sign := crypto.GetSignature(timestamp, ding.Secret, ding.Ticket)

	args := url.Values{}
	args.Set("accessKey", ding.Key)
	args.Set("suiteTicket", ding.Ticket)
	args.Set("timestamp", timestamp)
	args.Set("signature", sign)

	req := request.NewAuthInfo(ding.Key, corpId)
	res := &response.SuiteAccessToken{}
	if err := ding.Request(http.MethodPost, constant.GetAuthInfo, args, req, res); err != nil {
		return "", err
	}
	return res.Token, nil
}
