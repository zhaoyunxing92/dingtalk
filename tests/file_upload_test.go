package tests

import (
	"encoding/json"
	"github.com/zhaoyunxing92/dingtalk/model"
	"os"
	"testing"
)

func TestMediaUpload(t *testing.T) {

	file, err := os.Open("./../image/dingtalk.png")

	req := model.NewUploadFile("dingtalk.png", "image", file)
	media, err := dingTalk.MediaUpload(req)
	if err != nil {
		t.Fatal(err)
	}
	j, _ := json.Marshal(media)

	t.Log(string(j))
}
