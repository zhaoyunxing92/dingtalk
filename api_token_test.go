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
	"github.com/stretchr/testify/assert"
	"testing"
)

var client = NewClient(1244553273, "dingkjy4w80esdwgjuyo",
	"bDKa_nfJg3zYRsFrj-wTohTuoJCtxTEHaGmybYF9vgaVAZJOz-mICsLGStB288nW")

//var client = NewClient(1354379668, "ding4xqt9h4tb5hmovlq",
//	"rNzB2gWYLG5ZjVtnSCxk19nmJo1FQytEG5UsvzizIGLJYkwdfITsnjBs31AKQJdM").
//	Build()

func TestDingTalk_GetAccessToken(t *testing.T) {

	token, err := client.GetAccessToken()
	assert.Nil(t, err)
	assert.NotNil(t, token)
	t.Log(token)
}

func TestDingTalk_GetSuiteAccessToken(t *testing.T) {

	token, err := client.GetSuiteAccessToken()
	assert.Nil(t, err)
	assert.NotNil(t, token)
	t.Log(token)
}

func TestDingTalk_GetCorpAccessToken(t *testing.T) {

	ding := NewClient(10435002, "suitegqebx814dt8immqw",
		"AXjjwYhZ7Bwh1e8vlkg7pPQHUACwl8rSJWFma1taYMDLUjmIAtl9d9yAdTBg4K3m",
		WithTicket("60LVmSzP5zuDorEjISt2hc1imVeyKHcnTCUGUgslOBGpJXxyyQsbuBpyVhw6eunNAt3WXO0u9yfnXmcsskHNdN"),
		WithCorpId("ding3b1e912394011559bc961a6cb783455b"))

	token, err := ding.GetCorpAccessToken()
	assert.Nil(t, err)
	assert.NotNil(t, token)
	t.Log(token)
}

func TestDingTalk_GetSSOToken(t *testing.T) {

	token, err := client.GetSSOToken("dingc7c5220402493357f2c783f7214b6d69", "")

	assert.NotNil(t, err)
	assert.Nil(t, token)
	t.Log(token)
}

func TestDingTalk_GetJsApiTicket(t *testing.T) {

	ticket, err := client.GetJsApiTicket()

	assert.Nil(t, err)
	assert.NotNil(t, ticket)
	t.Log(ticket)
}
