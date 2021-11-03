package tests

import (
	"encoding/json"
	"github.com/zhaoyunxing92/dingtalk/v2/model"
	"os"
	"testing"
	"text/template"
)

func TestTemplate(t *testing.T) {
	str := `"当前应用id：{{.agentId}}"
"应用名称：{{.name}}"
`

	app := new(model.MicroApp)
	app.AgentId = 1020345059

	tmpl, err := template.New("文本渲染").Parse(str)
	if err != nil {
		t.Fatal(err)
	}
	args:=`{"agentId":1020345059,"name":"赵云兴"}`
	var obj map[string]interface{}
	err = json.Unmarshal([]byte(args), &obj)


	if err = tmpl.Execute(os.Stdout, obj); err != nil {
		t.Fatal(err)
	}

	//t.Log(wr.Write())

}
