package tests

import (
	"testing"
)

//测试应用列表
func TestDingTalkGetMicroAppList(t *testing.T) {

	apps, err := dingTalk.GetMicroAppList()

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("应用列表:%v", apps)
}

func TestDingTalkGetMicroAppVisibleScopes(t *testing.T) {
	scopes, err := dingTalk.GetMicroAppVisibleScopes(970282753)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("应用可视作用域:%v", scopes)
}

func TestDingTalkGetMicroAppByAgentId(t *testing.T) {
	app, err := dingTalk.GetMicroAppByAgentId(970282753)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", app)
}
