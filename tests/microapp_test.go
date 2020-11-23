package tests

import (
	"testing"
)

func TestDingTalk_GetMicroAppList(t *testing.T) {

	apps, err := dingTalk.GetMicroAppList()

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("应用列表:%v", apps)
}

func TestDingTalk_GetMicroAppVisibleScopes(t *testing.T) {
	scopes, err := dingTalk.GetMicroAppVisibleScopes(970282753)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("应用可视作用域:%v", scopes)
}
