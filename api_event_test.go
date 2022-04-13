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

	"github.com/zhaoyunxing92/dingtalk/v2/crypto"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
)

func TestDingTalk_RegisterEvent(t *testing.T) {

	token := crypto.RandomString(8)
	secret := crypto.RandomString(43)
	event := request.NewRegisterEvent("https://colobu.com", token, secret, []string{"user_add_org"})

	res, err := client.RegisterEvent(event)

	assert.NotNil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetFailedRegisterEvent(t *testing.T) {

	res, err := client.GetRegisterFailedEvent()

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_UpdateRegisterEvent(t *testing.T) {
	token := crypto.RandomString(8)
	secret := crypto.RandomString(43)
	event := request.NewRegisterEvent("https://colobu.com", token, secret, []string{"user_add_org"})

	res, err := client.UpdateRegisterEvent(event)

	assert.NotNil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetRegisterEvent(t *testing.T) {

	res, err := client.GetRegisterEvent()

	assert.Nil(t, err)
	assert.NotNil(t, res)

}

func TestDingTalk_DeleteRegisterEvent(t *testing.T) {

	res, err := client.DeleteRegisterEvent()

	assert.Nil(t, err)
	assert.NotNil(t, res)

}
