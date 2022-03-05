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
	"strings"
)

// LabelConfModify 修改角色或者角色组
type LabelConfModify struct {
	LabelConfAdd
	// 修改后的值
	PostLabels []PostLabel `json:"PostLabelList"`

	// 修改前的值
	PreLabels []PreLabel `json:"PreLabelList"`
}

// PostLabel 修改后
type PostLabel struct {
	Hidden bool `json:"hidden"`

	Name string `json:"name"`

	Id int `json:"id"`

	// 扩展信息
	Ext interface{} `json:"extraInfo"`
}

// PreLabel 修改前
type PreLabel struct {
	Hidden bool `json:"hidden"`

	Name string `json:"name"`

	Id int `json:"id"`

	// 扩展信息
	Ext interface{} `json:"extraInfo"`

	Deleted bool `json:"deleted"`

	Color int `json:"color"`

	Level int `json:"level"`

	Scope int `json:"scope"`

	ParentId int `json:"parentId"`
}

func (label *PreLabel) UnmarshalJSON(b []byte) (err error) {
	str := strings.Trim(string(b), "\"")
	str = strings.ReplaceAll(str, "\\", "")

	type TempPreLabel PreLabel

	ot := struct {
		*TempPreLabel // 避免直接嵌套PostLabel进入死循环
	}{
		TempPreLabel: (*TempPreLabel)(label),
	}
	if err = json.Unmarshal([]byte(str), &ot); err != nil {
		return err
	}
	return
}

// UnmarshalJSON "{\"hidden\":false,\"name\":\"权限组修改01\",\"id\":2393887343,\"extraInfo\":{}}"
func (label *PostLabel) UnmarshalJSON(b []byte) (err error) {
	str := strings.Trim(string(b), "\"")
	str = strings.ReplaceAll(str, "\\", "")

	type TempPostLabel PostLabel

	ot := struct {
		*TempPostLabel // 避免直接嵌套PostLabel进入死循环
	}{
		TempPostLabel: (*TempPostLabel)(label),
	}
	if err = json.Unmarshal([]byte(str), &ot); err != nil {
		return err
	}
	return
}
