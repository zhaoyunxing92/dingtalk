package model

//翻译返回
type AiResponse struct {
	Response
	Result string `json:"result"` //翻译结果字符串
}

//OcrStructuredResult
type OcrStructuredResult struct {
	Height         int    `json:"height"`          //旋转后图片高度。
	Width          int    `json:"width"`           //旋转后图片宽度。
	Angle          int    `json:"angle"`           //旋转度。
	Data           string `json:"data"`            //图片识别内容json字符串。
	OriginalHeight int    `json:"original_height"` //图片识别内容json字符串。
	OriginalWidth  int    `json:"original_width"`  //图片识别内容json字符串。
}

type OcrStructuredResponse struct {
	Response
	Result OcrStructuredResult `json:"result"` //翻译结果字符串
}
