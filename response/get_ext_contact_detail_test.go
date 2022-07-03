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
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExtContactDetail(t *testing.T) {
	str := `{
        "errcode":0,
        "result":{
                "share_user_ids":[
                        "66220007745510"
                ],
                "follower_user_id":"manager4220",
                "address":"北京",
                "share_dept_ids":[],
                "label_ids":[
                        1561077317
                ],
                "share_dept_ids": [
                     1561077317
                ],
                "mobile":"136***47616",
                "remark":"技术负责人",
                "title":"CFO",
                "userid":"011250026469774889",
                "company_name":"钉钉",
                "name":"王经理",
                "state_code":"86",
                "email":"zhang@example.com"
        },
        "request_id":"xrudojam4g8f"
}`
	res := &GetExtContactDetail{}
	err := json.Unmarshal([]byte(str), res)

	assert.Nil(t, err)
	assert.Equal(t, res.Code, 0)
	assert.Equal(t, res.RequestId, "xrudojam4g8f")
	assert.Equal(t, res.Result.ShareUser[0], "66220007745510")
	assert.Equal(t, res.Result.FollowerUser, "manager4220")
	assert.Equal(t, res.Result.Address, "北京")
	assert.Equal(t, len(res.Result.ShareDept), 1)
	assert.Equal(t, len(res.Result.Labels), 1)
	assert.Equal(t, res.Result.Labels[0], 1561077317)
	assert.Equal(t, res.Result.ShareDept[0], 1561077317)
	assert.Equal(t, res.Result.Mobile, "136***47616")
	assert.Equal(t, res.Result.Remark, "技术负责人")
	assert.Equal(t, res.Result.Title, "CFO")
	assert.Equal(t, res.Result.UserId, "011250026469774889")
	assert.Equal(t, res.Result.CompanyName, "钉钉")
	assert.Equal(t, res.Result.Name, "王经理")
	assert.Equal(t, res.Result.StateCode, "86")
	assert.Equal(t, res.Result.Email, "zhang@example.com")
}
