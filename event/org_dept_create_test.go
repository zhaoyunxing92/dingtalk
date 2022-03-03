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

func TestOrgDeptCreate(t *testing.T) {
	str := `{
   "CorpId": "dingc7c5220402493357f2c783f7214b6d69",
   "EventType": "org_dept_create",
   "DeptId": [
       581208264
   ],
   "TimeStamp": "1640669222083"
}`
	dept := &OrgDeptCreate{}

	err := json.Unmarshal([]byte(str), dept)

	assert.Nil(t, err)
	assert.Equal(t, dept.CorpId, "dingc7c5220402493357f2c783f7214b6d69")
	assert.Equal(t, dept.EventType, "org_dept_create")
	assert.Equal(t, dept.TimeStamp, 1640669222083)
	assert.Equal(t, len(dept.DeptIds), 1)
}
