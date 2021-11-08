package tests

import (
	"encoding/json"
	"testing"
)

//TestTranslate:测试翻译
func TestTranslate(t *testing.T) {
	//cidQm8R6iRg2CcePBHqlUU3xQ==
	rsp, err := dingTalk.Translate("golang 是世界上最好的语言", "zh", "en")
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}

//TestTranslate:测试翻译
func TestOcrRecognize(t *testing.T) {
	//cidQm8R6iRg2CcePBHqlUU3xQ==
	imageUrl := "http://imgsrc.baidu.com/forum/w%3D580/sign=10f06b6dd562853592e0d229a0ee76f2/edf00b46f21fbe09e0c3d28b6b600c338644ad4f.jpg"
	rsp, err := dingTalk.OcrRecognize("car_no", imageUrl)
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(rsp)
	t.Log(string(js))
}
