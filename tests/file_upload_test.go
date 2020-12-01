package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/lihongchen/dingtalk/model"
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
