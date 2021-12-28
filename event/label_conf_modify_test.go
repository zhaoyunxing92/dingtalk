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

func TestLabelConfModify(t *testing.T) {
	str := `{"PostLabelList":["{\"hidden\":false,\"name\":\"角色编辑测试\",\"id\":2393865715,\"extraInfo\":{}}"],"CorpId":"dingc7c5220402493357f2c783f7214b6d69","EventType":"label_conf_modify","PreLabelList":["{\"deleted\":false,\"color\":-11687445,\"hidden\":false,\"level\":1,\"scope\":1,\"name\":\"角色编辑\",\"id\":2393865715,\"parentId\":2393887343,\"extraInfo\":{}}"],"LabelIdList":[2393865715],"scope":"1","TimeStamp":"1640681578721"}`

	label := &LabelConfModify{}

	err := json.Unmarshal([]byte(str), label)

	assert.Nil(t, err)
}

//"{\"hidden\":false,\"name\":\"权限组修改01\",\"id\":2393887343,\"extraInfo\":{}}"
func TestPostLabel(t *testing.T) {
	str := "{\"hidden\":false,\"name\":\"权限组修改01\",\"id\":2393887343,\"extraInfo\":{}}"

	label := &PostLabel{}

	err := json.Unmarshal([]byte(str), label)

	assert.Nil(t, err)
}

// TestLabelConfModify_Role 角色修改
func TestLabelConfModify_Role(t *testing.T) {
	str := `{
    "PostLabelList": [
        {
            "hidden": false,
            "name": "角色编辑测试",
            "id": 2393865715,
            "extraInfo": {
                
            }
        }
    ],
    "CorpId": "dingc7c5220402493357f2c783f7214b6d69",
    "EventType": "label_conf_modify",
    "PreLabelList": [
        {
            "deleted": false,
            "color": -11687445,
            "hidden": false,
            "level": 1,
            "scope": 1,
            "name": "角色编辑",
            "id": 2393865715,
            "parentId": 2393887343,
            "extraInfo": {
                
            }
        }
    ],
    "LabelIdList": [
        2393865715
    ],
    "scope": "1",
    "TimeStamp": "1640681578721"
}`

	label := &LabelConfModify{}

	err := json.Unmarshal([]byte(str), label)

	assert.Nil(t, err)
	assert.Equal(t, label.TimeStamp, 1640681578721)
	assert.Equal(t, label.Scope, "1")

	assert.Equal(t, len(label.LabelIds), 1)
	assert.Equal(t, len(label.PreLabels), 1)
	assert.Equal(t, len(label.PostLabels), 1)

	assert.Equal(t, label.PreLabels[0].Deleted, false)
	assert.Equal(t, label.PreLabels[0].Color, -11687445)
	assert.Equal(t, label.PreLabels[0].Hidden, false)
	assert.Equal(t, label.PreLabels[0].Level, 1)
	assert.Equal(t, label.PreLabels[0].Scope, 1)
	assert.Equal(t, label.PreLabels[0].Name, "角色编辑")
	assert.Equal(t, label.PreLabels[0].Id, 2393865715)
	assert.Equal(t, label.PreLabels[0].ParentId, 2393887343)

	assert.Equal(t, label.PostLabels[0].Hidden, false)
	assert.Equal(t, label.PostLabels[0].Name, "角色编辑测试")
	assert.Equal(t, label.PostLabels[0].Id, 2393865715)

	assert.Equal(t, label.EventType, "label_conf_modify")
	assert.Equal(t, label.CorpId, "dingc7c5220402493357f2c783f7214b6d69")
}

// TestLabelConfModify_Role_Group 角色组修改
func TestLabelConfModify_Role_Group(t *testing.T) {
	str := `{
    "PostLabelList": [
        {
            "hidden": false,
            "name": "权限组修改",
            "id": 2393887343,
            "extraInfo": {
                
            }
        }
    ],
    "CorpId": "dingc7c5220402493357f2c783f7214b6d69",
    "EventType": "label_conf_modify",
    "PreLabelList": [
        {
            "deleted": false,
            "color": -11687445,
            "hidden": false,
            "level": 0,
            "scope": 1,
            "name": "权限组修改01",
            "id": 2393887343,
            "parentId": -1,
            "extraInfo": {
                
            }
        }
    ],
    "LabelIdList": [
        2393887343
    ],
    "scope": "1",
    "TimeStamp": "1640682321572"
}`

	label := &LabelConfModify{}

	err := json.Unmarshal([]byte(str), label)

	assert.Nil(t, err)
	assert.Equal(t, label.TimeStamp, 1640682321572)
	assert.Equal(t, label.Scope, "1")

	assert.Equal(t, len(label.LabelIds), 1)
	assert.Equal(t, len(label.PreLabels), 1)
	assert.Equal(t, len(label.PostLabels), 1)

	assert.Equal(t, label.PreLabels[0].Deleted, false)
	assert.Equal(t, label.PreLabels[0].Color, -11687445)
	assert.Equal(t, label.PreLabels[0].Hidden, false)
	assert.Equal(t, label.PreLabels[0].Level, 0)
	assert.Equal(t, label.PreLabels[0].Scope, 1)
	assert.Equal(t, label.PreLabels[0].Name, "权限组修改01")
	assert.Equal(t, label.PreLabels[0].Id, 2393887343)
	assert.Equal(t, label.PreLabels[0].ParentId, -1)

	assert.Equal(t, label.PostLabels[0].Hidden, false)
	assert.Equal(t, label.PostLabels[0].Name, "权限组修改")
	assert.Equal(t, label.PostLabels[0].Id, 2393887343)

	assert.Equal(t, label.EventType, "label_conf_modify")
	assert.Equal(t, label.CorpId, "dingc7c5220402493357f2c783f7214b6d69")
}
