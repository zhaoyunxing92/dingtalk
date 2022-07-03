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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeptDetailUserInfo(t *testing.T) {

	str := `{
    "errcode": 0,
    "result": {
        "next_cursor": 10,
        "has_more": true,
        "list": [
            {
                "leader": true,
                "extension": "{\"爱好\":\"旅游\",\"年龄\":\"24\"}",
                "unionid": "z21HjQliSzpw0YWCNxmii6u2Os62cZ62iSZ",
                "boss": true,
                "exclusive_account": false,
                "admin": true,
                "remark": "备注备注",
                "title": "技术总监",
                "hired_date": 1597573616828,
                "userid": "zhangsan",
                "work_place": "未来park",
                "job_number": "4",
                "dept_id_list":[2,3,4],
                "email": "test@xxx.com",
                "dept_order": 1,
                "mobile": "18513027676",
                "active": true,
                "telephone": "010-86123456-2345",
                "avatar": "xxx",
                "hide_mobile": false,
                "org_email": "test@xxx.com",
                "name": "张三",
                "state_code": "86"
            }
        ]
    },
    "errmsg": "ok"
}`
	res := &DeptDetailUserInfo{}
	err := json.Unmarshal([]byte(str), res)

	assert.Nil(t, err)
	assert.Equal(t, res.Code, 0)
	assert.Equal(t, res.Msg, "ok")
	assert.Equal(t, res.Page.NextCursor, 10)
	assert.Equal(t, res.Page.HasMore, true)
	assert.Equal(t, res.Page.List[0].Leader, true)
	assert.Equal(t, res.Page.List[0].UnionId, "z21HjQliSzpw0YWCNxmii6u2Os62cZ62iSZ")
	assert.Equal(t, res.Page.List[0].Boss, true)
	assert.Equal(t, res.Page.List[0].ExclusiveAccount, false)
	assert.Equal(t, res.Page.List[0].Admin, true)
	assert.Equal(t, res.Page.List[0].Remark, "备注备注")
	assert.Equal(t, res.Page.List[0].Title, "技术总监")
	assert.Equal(t, res.Page.List[0].HiredDate, 1597573616828)
	assert.Equal(t, res.Page.List[0].UserId, "zhangsan")
	assert.Equal(t, res.Page.List[0].WorkPlace, "未来park")
	assert.Equal(t, res.Page.List[0].JobNumber, "4")
	assert.Equal(t, res.Page.List[0].Email, "test@xxx.com")
	assert.Equal(t, res.Page.List[0].DeptOrder, 1)
	assert.Equal(t, res.Page.List[0].Mobile, "18513027676")
	assert.Equal(t, res.Page.List[0].Active, true)
	assert.Equal(t, res.Page.List[0].Telephone, "010-86123456-2345")
	assert.Equal(t, res.Page.List[0].Avatar, "xxx")
	assert.Equal(t, res.Page.List[0].HideMobile, false)
	assert.Equal(t, res.Page.List[0].OrgEmail, "test@xxx.com")
	assert.Equal(t, res.Page.List[0].Name, "张三")
	assert.Equal(t, res.Page.List[0].StateCode, "86")
	assert.Equal(t, res.Page.List[0].DeptIds[0], 2)
	assert.Equal(t, res.Page.List[0].DeptIds[1], 3)
	assert.Equal(t, res.Page.List[0].DeptIds[2], 4)

	du := DeptUser{}

	du.Name = "张三"
	du.Mobile = "18513027676"
	du.DeptIds = []int{2, 3, 4}

	res.Page.List[0] = du
	res.Page.NextCursor = 10
	res.Page.HasMore = false

	marshal, err := json.Marshal(res)

	assert.Nil(t, err)
	fmt.Println(string(marshal))
}
