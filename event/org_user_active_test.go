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

func TestOrgUserActive(t *testing.T) {
	str := `{
  "CorpId": "dinge8a5xxxx",
  "EventType": "user_active_org",
  "UserId": [
    "0119296xxx"
  ],
  "TimeStamp": "1608017258073"
}`
	user := &OrgUserActive{}

	err := json.Unmarshal([]byte(str), user)

	assert.Nil(t, err)
	assert.Equal(t, user.EventType, "user_active_org")
	assert.Equal(t, user.TimeStamp, 1608017258073)
	assert.Equal(t, len(user.UserIds), 1)
	assert.Equal(t, user.CorpId, "dinge8a5xxxx")
}
