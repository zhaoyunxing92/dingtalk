/*
 * Copyright 2020 zhaoyunxing.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dingtalk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDingTalk_GetAuthInfo(t *testing.T) {
	t.Skip()
	res, err := isv.GetAuthInfo("ding3b1e912394011559bc961a6cb783455b")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.AuthCorpInfo.CorpId, "ding3b1e912394011559bc961a6cb783455b")
}

func TestDingTalk_GetAgentInfo(t *testing.T) {
	t.Skip()
	res, err := isv.GetAgentInfo(1332307896, "ding3b1e912394011559bc961a6cb783455b")

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetUnactiveCorp(t *testing.T) {
	t.Skip()
	res, err := isv.GetUnactiveCorp(45829)

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_ReauthCorp(t *testing.T) {
	t.Skip()
	res, err := isv.ReauthCorp(45829, "ding3b1e912394011559bc961a6cb783455b")

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetCorpPermanentCode(t *testing.T) {
	t.Skip()
	res, err := isv.GetCorpPermanentCode("3aa3572685c93be1a0b2a6b3dba88d4f")

	assert.Nil(t, err)
	assert.NotNil(t, res)
}
