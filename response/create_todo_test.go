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

func TestCreateTodo(t *testing.T) {
	str := `{
  "id" : "OPJpwtwPVNGIFKURjrzd",
  "subject" : "接入钉钉待办",
  "description" : "应用可以调用该接口发起一个钉钉待办任务",
  "startTime" : 1617675000000,
  "dueTime" : 1617675100000,
  "finishTime" : 1617675200000,
  "done" : false,
  "executorIds" : [ "PUoiinWIpa2yH2ymhiiGiP6g" ],
  "participantIds" : [ "PUoiinWIpa2yH2ymhiiGiP6g" ],
  "detailUrl" : {
    "pcUrl" : "https://www.dingtalk.com",
    "appUrl" : "https://www.dingtalk.com"
  },
  "source" : "isv_dingtalkTodo",
  "sourceId" : "isv_dingtalkTodo1",
  "createdTime" : 1617675200000,
  "modifiedTime" : 1617675200000,
  "creatorId" : "PUoiinWIpa2yH2ymhiiGiP6g",
  "modifierId" : "PUoiinWIpa2yH2ymhiiGiP6g",
  "bizTag" : "isv_dingtalkTodo",
  "requestId" : "PUoiinWIpa2yH2ymhiiGiP6g",
  "isOnlyShowExecutor" : true,
  "priority" : 20,
  "notifyConfigs" : { 
    "dingNotify":"ding"
   }
}`
	res := &CreateTodo{}
	err := json.Unmarshal([]byte(str), res)

	assert.Nil(t, err)
	assert.Equal(t, res.Id, "OPJpwtwPVNGIFKURjrzd")
	assert.Equal(t, res.Subject, "接入钉钉待办")
	assert.Equal(t, res.Desc, "应用可以调用该接口发起一个钉钉待办任务")
	assert.Equal(t, res.StartTime, 1617675000000)
	assert.Equal(t, res.DueTime, 1617675100000)
	assert.Equal(t, res.FinishTime, 1617675200000)
	assert.Equal(t, res.Done, false)
	assert.Equal(t, res.Executors[0], "PUoiinWIpa2yH2ymhiiGiP6g")
	assert.Equal(t, res.Participants[0], "PUoiinWIpa2yH2ymhiiGiP6g")
	assert.Equal(t, res.Urls.Pc, "https://www.dingtalk.com")
	assert.Equal(t, res.Urls.App, "https://www.dingtalk.com")
	assert.Equal(t, res.Source, "isv_dingtalkTodo")
	assert.Equal(t, res.SourceId, "isv_dingtalkTodo1")
	assert.Equal(t, res.CreatedTime, 1617675200000)
	assert.Equal(t, res.ModifiedTime, 1617675200000)
	assert.Equal(t, res.CreatorId, "PUoiinWIpa2yH2ymhiiGiP6g")
	assert.Equal(t, res.ModifierId, "PUoiinWIpa2yH2ymhiiGiP6g")
	assert.Equal(t, res.BizTag, "isv_dingtalkTodo")
	assert.Equal(t, res.RequestId, "PUoiinWIpa2yH2ymhiiGiP6g")
	assert.Equal(t, res.OnlyShowExecutor, true)
	assert.Equal(t, res.Priority, 20)
	assert.Equal(t, res.NotifyConfigs.Ding, "ding")

}
