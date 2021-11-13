/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSendTemplateMessage(t *testing.T) {
	data := make(map[string]string, 2)
	data["name"] = "merge"
	data["age"] = "18"

	msg := NewSendTemplateMessage(1332307896, "80424a44cb444c53a071aae34c0fd140").
		SetUserIds("manager164", "manager164").
		SetMessage(data).
		SetData("msg", "发送模板消息").
		SetData("name", "模板消息").
		Build()

	assert.NotNil(t, msg)
	assert.Equal(t, len(msg.UserIds), 1)
	assert.Equal(t, len(msg.Data), 3)
	assert.Equal(t, msg.Data["name"], "模板消息")
}
