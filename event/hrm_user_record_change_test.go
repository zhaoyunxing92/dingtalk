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

func TestHrmUserRecordChange(t *testing.T) {
	str := `{
   "actionType": "userInfoChange",
   "EventType": "hrm_user_record_change",
   "staffId": "manager164"
}`
	hrm := &HrmUserRecordChange{}

	err := json.Unmarshal([]byte(str), hrm)

	assert.Nil(t, err)
	assert.Equal(t, hrm.ActionType, "userInfoChange")
	assert.Equal(t, hrm.EventType, "hrm_user_record_change")
	assert.Equal(t, hrm.StaffId, "manager164")
}
