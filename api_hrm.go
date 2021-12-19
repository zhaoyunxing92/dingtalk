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
