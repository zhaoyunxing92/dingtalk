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

func TestBpmTaskChange_Start(t *testing.T) {
	str := `{
    "processInstanceId": "60406628-b642-4740-81e7-68cc67ebaa45",
    "corpId": "dingc7c5220402493357f2c783f7214b6d69",
    "EventType": "bpms_task_change",
    "businessId": "202112281713000090163",
    "title": "赵云提交的审批任务",
    "type": "start",
    "createTime": 1640682790000,
    "processCode": "PROC-5CB63684-096E-40DD-B326-CF60691C507A",
    "bizCategoryId": "",
    "staffId": "manager164",
    "taskId": 72062030325
}`

	bpm := &BpmTaskChange{}

	err := json.Unmarshal([]byte(str), bpm)

	assert.Nil(t, err)
	assert.Equal(t, bpm.Type, "start")
	assert.Equal(t, bpm.BizCategoryId, "")
	assert.Equal(t, bpm.StaffId, "manager164")
	assert.Equal(t, bpm.Title, "赵云提交的审批任务")
	assert.Equal(t, bpm.CreateTime, 1640682790000)
	assert.Equal(t, bpm.EventType, "bpms_task_change")
	assert.Equal(t, bpm.BusinessId, "202112281713000090163")
	assert.Equal(t, bpm.CorpId, "dingc7c5220402493357f2c783f7214b6d69")
	assert.Equal(t, bpm.ProcessCode, "PROC-5CB63684-096E-40DD-B326-CF60691C507A")
	assert.Equal(t, bpm.ProcessInstanceId, "60406628-b642-4740-81e7-68cc67ebaa45")
}

func TestBpmTaskChange_Refuse(t *testing.T) {
	str := `{
    "processInstanceId": "45437459-7a8d-4af8-b16b-b535c80f2756",
    "finishTime": 1640684616000,
    "corpId": "dingc7c5220402493357f2c783f7214b6d69",
    "EventType": "bpms_task_change",
    "businessId": "202112281743000242367",
    "title": "赵云提交的审批任务",
    "type": "finish",
    "result": "refuse",
    "createTime": 1640684605000,
    "processCod": "5CB63684-096E-40DD-B326-CF60691C507A",
    "bizCategoryId": "",
    "staffId": "manager164",
    "taskId": 72025621907
}`

	bpm := &BpmTaskChange{}

	err := json.Unmarshal([]byte(str), bpm)

	assert.Nil(t, err)
	assert.Equal(t, bpm.Type, "finish")
	assert.Equal(t, bpm.Result, "refuse")
	assert.Equal(t, bpm.BizCategoryId, "")
	assert.Equal(t, bpm.StaffId, "manager164")
	assert.Equal(t, bpm.Title, "赵云提交的审批任务")
	assert.Equal(t, bpm.TaskId, 72025621907)
	assert.Equal(t, bpm.CreateTime, 1640684605000)
	assert.Equal(t, bpm.FinishTime, 1640684616000)
	assert.Equal(t, bpm.EventType, "bpms_task_change")
	assert.Equal(t, bpm.BusinessId, "202112281713000090163")
	assert.Equal(t, bpm.CorpId, "dingc7c5220402493357f2c783f7214b6d69")
	assert.Equal(t, bpm.ProcessCode, "PROC-5CB63684-096E-40DD-B326-CF60691C507A")
	assert.Equal(t, bpm.ProcessInstanceId, "60406628-b642-4740-81e7-68cc67ebaa45")
}
