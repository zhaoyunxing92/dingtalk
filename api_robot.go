/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dingtalk

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zhaoyunxing92/dingtalk/v2/domain"
	"net/http"
	"time"
)

type Robot struct {
	Token    string
	client   *http.Client
	validate *validator.Validate //参数校验
	trans    translator.Translator
}

//https://oapi.dingtalk.com/robot/send?access_token=f00184626027aafcd6b4fe07b90ec11250d78b8735282185256a0c72186a5f49
func NewRobot(token string) Robot {
	validate := validator.New()
	uni := translator.New(en.New(), zh.New())
	trans, _ := uni.GetTranslator("zh")
	_ = zh_trans.RegisterDefaultTranslations(validate, trans)

	return Robot{token, &http.Client{Timeout: 10 * time.Second}, validate, trans}
}

//SendRobotMsg:机器人发送消息
func (robot *Robot) SendRobotMsg(req domain.Request) (res domain.Response, err error) {

	if err = req.Validate(robot.validate, robot.trans); err != nil {
		return res, err
	}
	err = robot.send(req, &res)
	return res, err
}
