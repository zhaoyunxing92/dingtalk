package dingtalk

import (
	"github.com/pkg/errors"
	"github.com/zhaoyunxing92/dingtalk/v2/cache"
	"net/http"
	"strings"
	"time"
)

type dingTalk struct {
	// 企业内部应用对应:AgentId，如果是应套件:SuiteId
	Id int `json:"id,omitempty" validate:"required"`

	// 企业内部应用对应:AppKey，套件对应:SuiteKey
	Key string `json:"key,omitempty" validate:"required"`

	//企业内部对应:AppSecret，套件对应:SuiteSecret
	Secret string `json:"secret,omitempty" validate:"required"`

	// isv 钉钉开放平台会向应用的回调URL推送的suite_ticket（约5个小时推送一次）
	Ticket string `json:"ticket,omitempty"`

	//授权企业的id
	CorpId string `json:"corpId"`

	Client *http.Client

	Cache cache.Cache
}

//isv 是否isv
func (ding *dingTalk) isv() bool {
	return len(ding.Ticket) > 0 && len(ding.CorpId) > 0
}

// builder builder for dingtalk
type builder struct {
	ding *dingTalk
}

// NewClient new DingTalkBuilder
func NewClient(id int, key, secret string) *builder {
	return &builder{ding: &dingTalk{Id: id, Key: key, Secret: secret}}
}

//SetTicket set app ticket
func (b *builder) SetTicket(ticket string) *builder {
	if len(ticket) <= 0 {
		panic(errors.New("ticket is invalid"))
	}
	b.ding.Ticket = ticket
	return b
}

//SetCorpId set app ticket
func (b *builder) SetCorpId(corpId string) *builder {
	if len(corpId) <= 0 {
		panic(errors.New("corpId is invalid"))
	}
	b.ding.CorpId = corpId
	return b
}

//Build return dingtalk
func (b *builder) Build() *dingTalk {

	if err := validate(b); err != nil {
		panic(err)
	}
	ding := b.ding
	key := ding.Key
	//判断是否isv
	if ding.isv() {
		ding.Cache = cache.NewFileCache(strings.Join([]string{".token", "isv", key}, "/"), ding.CorpId)
	} else {
		ding.Cache = cache.NewFileCache(strings.Join([]string{".token", "corp"}, "/"), key)
	}
	ding.Client = &http.Client{Timeout: 10 * time.Second}
	return b.ding
}
