package dingtalk_go

import (
	"testing"
)

func TestDingTalk_GetToken(t *testing.T) {
	talk, _ := NewDingTalk("dingdv9vbx9mcrj18g1n", "bDBJd67ct1Ik0GFqhNNWH4Lbo4aqZGglaE1wJ3mnbG6ANRjOruuGzs6Z0glNEU63")

	token, err := talk.GetToken()

	t.Log(token)
	t.Log(err)
}
