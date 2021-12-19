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
	"github.com/zhaoyunxing92/dingtalk/v2/constant/employee"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

// GetHrmEmployee 获取在职员工列表
func (ding dingTalk) GetHrmEmployee(offset, size int, state employee.Status, status ...employee.Status) (res response.GetHrmEmployee,
	err error,
) {
	req := request.NewGetHrmEmployee(offset, size, append(status, state))
	return res, ding.Request(http.MethodPost, constant.GetHrmEmployeeKey, nil, req, &res)
}
