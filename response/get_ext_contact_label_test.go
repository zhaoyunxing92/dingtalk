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

func TestGetExtContactLabel(t *testing.T) {
	str := `{
    "errcode": 0,
    "results": [
        {
            "color": -15220075,
            "name": "类型",
            "labels": [
                {
                    "name": "客户",
                    "id": 1561077310
                }
            ]
        }
    ],
    "request_id": "51jof1jt5f4t"
}`
	res := &GetExtContactLabel{}
	err := json.Unmarshal([]byte(str), res)

	assert.Nil(t, err)
	assert.Equal(t, res.Code, 0)
	assert.Equal(t, res.RequestId, "51jof1jt5f4t")
	assert.Equal(t, res.Results[0].Color, -15220075)
	assert.Equal(t, res.Results[0].Name, "类型")
	assert.Equal(t, res.Results[0].Labels[0].Name, "客户")
	assert.Equal(t, res.Results[0].Labels[0].Id, 1561077310)
}
