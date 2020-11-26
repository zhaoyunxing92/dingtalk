package tests

import "testing"

func TestCreatUser(t *testing.T) {
	//011755000243774889
	user, err := dingTalk.CreateUser("张三", "18513027676", 1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", user)
}
