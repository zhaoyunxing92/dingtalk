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
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/crypto"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// GetAuthInfo 获取企业授权信息
func (ding *dingTalk) GetAuthInfo(corpId string) (string, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	sign := crypto.GetSignature(timestamp, ding.Secret, ding.Ticket)

	args := url.Values{}
	args.Set("accessKey", ding.Key)
	args.Set("suiteTicket", ding.Ticket)
	args.Set("timestamp", timestamp)
	args.Set("signature", sign)

	req := request.NewAuthInfo(ding.Key, corpId)
	res := &response.SuiteAccessToken{}
	if err := ding.Request(http.MethodPost, constant.GetAuthInfo, args, req, res); err != nil {
		return "", err
	}
	return res.Token, nil
}
