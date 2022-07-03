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

package response

import (
	json "encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeGetUserInfo(t *testing.T) {
	str := `{
        "errcode": 0, 
        "result": {
                "associated_unionid": "N2o5U3axxxx", 
                "unionid": "gliiW0piiii02zBUjUxxxx", 
                "device_id": "12drtfxxxxx", 
                "sys_level": 1, 
                "name": "张xx", 
                "sys": true, 
                "userid": "userid123"
        }, 
        "errmsg": "ok"
}`
	res := &CodeGetUserInfo{}
	err := json.Unmarshal([]byte(str), res)

	assert.Nil(t, err)
	assert.Equal(t, res.Code, 0)
	assert.Equal(t, res.Msg, "ok")
	assert.Equal(t, res.UserInfo.AssociatedUnionId, "N2o5U3axxxx")
	assert.Equal(t, res.UserInfo.UnionId, "gliiW0piiii02zBUjUxxxx")
	assert.Equal(t, res.UserInfo.DeviceId, "12drtfxxxxx")
	assert.Equal(t, res.UserInfo.Level, 1)
	assert.Equal(t, res.UserInfo.Name, "张xx")
	assert.Equal(t, res.UserInfo.Admin, true)
	assert.Equal(t, res.UserInfo.UserId, "userid123")

}
