package dingtalk

import (
	"net/http"
)
import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

//GetIndustryDeptDetail 获取部门详情
func (ding *dingTalk) GetIndustryDeptDetail(deptId int) (rsp response.GetIndustryDeptDetail, err error) {

	return rsp, ding.Request(http.MethodPost, constant.GetIndustryDeptDetailKey, nil,
		request.NewIndustryDeptDetail(deptId), &rsp)
}

//GetIndustryDept 获取部门列表,行业根部门传1
func (ding *dingTalk) GetIndustryDept(deptId, offset, size int) (rsp response.GetIndustryDept, err error) {

	return rsp, ding.Request(http.MethodPost, constant.GetIndustryDeptKey, nil,
		request.NewIndustryDept(deptId, offset, size), &rsp)
}
