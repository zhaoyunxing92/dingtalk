package tests

import (
	"encoding/json"
	"github.com/zhaoyunxing92/dingtalk/domain"
	"io/ioutil"
	"testing"
)

func TestMediaUpload(t *testing.T) {

	file, err := ioutil.ReadFile("./../image/dingtalk.png")

	req := domain.NewUploadFile("image", file)
	media, err := dingTalk.MediaUpload(req)
	if err != nil {
		t.Fatal(err)
	}
	j, _ := json.Marshal(media)

	t.Log(string(j))
}
