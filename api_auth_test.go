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

func TestDingTalk_GetAuthInfo(t *testing.T) {

	ding := NewClient(10435002, "suitegqebx814dt8immqw",
		"AXjjwYhZ7Bwh1e8vlkg7pPQHUACwl8rSJWFma1taYMDLUjmIAtl9d9yAdTBg4K3m",
		WithTicket("60LVmSzP5zuDorEjISt2hc1imVeyKHcnTCUGUgslOBGpJXxyyQsbuBpyVhw6eunNAt3WXO0u9yfnXmcsskHNdN"),
		WithCorpId("ding3b1e912394011559bc961a6cb783455b"))

	res, err := ding.GetAuthInfo("ding3b1e912394011559bc961a6cb783455b")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.AuthCorpInfo.CorpId, "ding3b1e912394011559bc961a6cb783455b")
}

func TestDingTalk_GetAgentInfo(t *testing.T) {

	ding := NewClient(10435002, "suitegqebx814dt8immqw",
		"AXjjwYhZ7Bwh1e8vlkg7pPQHUACwl8rSJWFma1taYMDLUjmIAtl9d9yAdTBg4K3m",
		WithTicket("60LVmSzP5zuDorEjISt2hc1imVeyKHcnTCUGUgslOBGpJXxyyQsbuBpyVhw6eunNAt3WXO0u9yfnXmcsskHNdN"),
		WithCorpId("ding3b1e912394011559bc961a6cb783455b"))

	res, err := ding.GetAgentInfo(1332307896, "ding3b1e912394011559bc961a6cb783455b")

	assert.Nil(t, err)
	assert.NotNil(t, res)

}

func TestDingTalk_GetUnactiveCorp(t *testing.T) {

	ding := NewClient(10435002, "suitegqebx814dt8immqw",
		"AXjjwYhZ7Bwh1e8vlkg7pPQHUACwl8rSJWFma1taYMDLUjmIAtl9d9yAdTBg4K3m",
		WithTicket("60LVmSzP5zuDorEjISt2hc1imVeyKHcnTCUGUgslOBGpJXxyyQsbuBpyVhw6eunNAt3WXO0u9yfnXmcsskHNdN"),
		WithCorpId("ding3b1e912394011559bc961a6cb783455b"))

	res, err := ding.GetUnactiveCorp(45829)

	assert.Nil(t, err)
	assert.NotNil(t, res)

}

func TestDingTalk_ReauthCorp(t *testing.T) {

	ding := NewClient(10435002, "suitegqebx814dt8immqw",
		"AXjjwYhZ7Bwh1e8vlkg7pPQHUACwl8rSJWFma1taYMDLUjmIAtl9d9yAdTBg4K3m",
		WithTicket("60LVmSzP5zuDorEjISt2hc1imVeyKHcnTCUGUgslOBGpJXxyyQsbuBpyVhw6eunNAt3WXO0u9yfnXmcsskHNdN"),
		WithCorpId("ding3b1e912394011559bc961a6cb783455b"))

	res, err := ding.ReauthCorp(45829, "ding3b1e912394011559bc961a6cb783455b")

	assert.Nil(t, err)
	assert.NotNil(t, res)

}
