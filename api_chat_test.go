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
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"testing"
)

func TestDingTalk_CreateChat(t *testing.T) {
	res, err := client.CreateChat(
		request.NewCreatChat("dingtalk", "manager164",
			"manager164", "manager164", "manager164").
			Build())

	// chat79e713edce46ea72ea0dd60fb4f9d6f6
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetChatInfo(t *testing.T) {
	t.Skip()
	_, _ = client.GetChatInfo("chat8ff884ef696f5717678c6280edfdbbf1")

}

func TestDingTalk_UpdateChat(t *testing.T) {

	res, err := client.UpdateChat(
		request.NewUpdateChat("chat8ff884ef696f5717678c6280edfdbbf1").
			SetName("java").
			SetShowHistory(0).
			SetSearchable(1).
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetChatQRCode(t *testing.T) {

	res, err := client.GetChatQRCode("chat8ff884ef696f5717678c6280edfdbbf1", "manager164")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.NotNil(t, res.Url)
}
