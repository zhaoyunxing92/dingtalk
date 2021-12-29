package dingtalk

import (
	"net/http"

	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

// GetAttendanceGroups 批量获取考勤组详情
func (ding *dingTalk) GetAttendanceGroups(offset, size int) (rsp response.GetAttendanceGroup, err error) {
	return rsp, ding.Request(http.MethodPost, constant.GetAttendanceGroupsKey, nil,
		request.NewGetAttendanceGroup(offset, size), &rsp)
}
