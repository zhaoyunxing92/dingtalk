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

func TestCorpPermanentCode(t *testing.T) {
	str := `{
        "errcode":0,
        "auth_corp_info":{
                "corpid":"xxxxx",
                "corp_name":"name"
        },
        "errmsg":"ok",
        "permanent_code":"xxxx"
}`
	res := &CorpPermanentCode{}

	err := json.Unmarshal([]byte(str), res)

	assert.Nil(t, err)
	assert.Equal(t, res.Code, 0)
	assert.Equal(t, res.Msg, "ok")
	assert.Equal(t, res.PermanentCode, "xxxx")
	assert.Equal(t, res.CorpInfo.CorpId, "xxxxx")
	assert.Equal(t, res.CorpInfo.CorpName, "name")
}
