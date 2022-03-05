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

package event

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabelConfDel(t *testing.T) {
	str := `{
    "CorpId": "dingc7c5220402493357f2c783f7214b6d69",
    "EventType": "label_conf_del",
    "LabelIdList": [
        2393849326
    ],
    "scope": "1",
    "TimeStamp": "1640673879734"
}`

	label := &LabelConfDel{}

	err := json.Unmarshal([]byte(str), label)

	assert.Nil(t, err)
	assert.Equal(t, label.Scope, "1")
	assert.Equal(t, len(label.LabelIds), 1)
	assert.Equal(t, label.TimeStamp, 1640673879734)
	assert.Equal(t, label.EventType, "label_conf_del")
	assert.Equal(t, label.CorpId, "dingc7c5220402493357f2c783f7214b6d69")
}
