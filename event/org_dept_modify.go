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

// OrgDeptModify 通讯录企业部门修改
// {
//    "CorpId": "dingc7c5220402493357f2c783f7214b6d69",
//    "EventType": "org_dept_modify",
//    "DeptId": [
//        581208264
//    ],
//    "TimeStamp": "1640669396434"
//}
type OrgDeptModify struct {
	OrgDeptCreate
}
