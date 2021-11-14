package message

import (
	"testing"
)

func TestNewCardMessage(t *testing.T) {
	content := `# 这是支持markdown的文本
## 标题2
* 列表1

![alt 啊](https://img.alicdn.com/tps/TB1XLjqNVXXXXc4XVXXXXXXXXXX-170-64.png)`

	msg := NewCardMessage(content)

	msg.Title = "是透出到会话列表和通知的文案"
	msg.SingleTitle = "查看详情"
	msg.SingleUrl = "https://open.dingtalk.com"

	msg.BtnOrientation = "1"
	msg.CardButtons = []CardButton{
		{"一个按钮", "https://www.taobao.com"},
		{"两个按钮", "https://www.tmall.com"},
	}

	t.Log(msg.String())

}
