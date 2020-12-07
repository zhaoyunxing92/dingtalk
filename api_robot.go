package dingtalk

import (
	"github.com/zhaoyunxing92/dingtalk/model"
	"net/http"
	"time"
)

type Robot struct {
	Token  string
	client *http.Client
}
//https://oapi.dingtalk.com/robot/send?access_token=f00184626027aafcd6b4fe07b90ec11250d78b8735282185256a0c72186a5f49
func NewRobot(token string) Robot {
	return Robot{token, &http.Client{Timeout: 10 * time.Second}}
}

//SendRobotMsg:机器人发送消息
func (robot *Robot) SendRobotMsg(req model.Request) (res model.Response, err error) {
	err = robot.send(req, &res)
	return res, err
}
