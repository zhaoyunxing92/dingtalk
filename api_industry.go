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

package dingtalk

import (
	"net/http"

	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

// GetIndustryDeptDetail 获取部门详情
func (ding *dingTalk) GetIndustryDeptDetail(deptId int) (rsp response.GetIndustryDeptDetail, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetIndustryDeptDetailKey, nil,
		request.NewIndustryDeptDetail(deptId), &rsp)
}

// GetIndustryDept 获取部门列表,行业根部门传1
func (ding *dingTalk) GetIndustryDept(deptId, offset, size int) (rsp response.GetIndustryDept, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetIndustryDeptKey, nil,
		request.NewIndustryDept(deptId, offset, size), &rsp)
}
