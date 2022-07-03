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

func TestGetParentIdsByUserId(t *testing.T) {
	str := `{
    "errcode": 0,
    "errmsg":"ok",
    "result": {
        "parent_list": [
            {
                "parent_dept_id_list": [
                    3995,
                    1
                ]
            }
        ]
    },
    "request_id": "nfe328zf67ec"
}`
	res := &GetParentIdsByUserId{}
	err := json.Unmarshal([]byte(str), res)

	assert.Nil(t, err)
	assert.Equal(t, res.Code, 0)
	assert.Equal(t, res.Msg, "ok")
	assert.Equal(t, res.Result.Parents[0].ParentIds[0], 3995)
	assert.Equal(t, res.Result.Parents[0].ParentIds[1], 1)
}
