package response

type FailedRegisterEvent struct {
	Response

	// 是否还有推送失败的变更事件，若为true，则表示还有未回调的事件
	HasMore bool `json:"has_more"`

	// 推送失败的事件列表，一次最多200个
	FailedList []struct {
		//
		UserAddOrg struct {
			Userid []string `json:"userid"`
			Corpid string   `json:"corpid"`
		} `json:"user_add_org,omitempty"`

		CallBackTag        string `json:"call_back_tag"`
		EventTime          int64  `json:"event_time"`
		BpmsInstanceChange struct {
			BpmsCallBackData struct {
			} `json:"bpmsCallBackData"`
			Corpid string `json:"corpid"`
		} `json:"bpms_instance_change,omitempty"`
		LabelConfAdd struct {
			RoleLabelChange struct {
			} `json:"roleLabelChange"`
			Corpid string `json:"corpid"`
		} `json:"label_conf_add,omitempty"`
	} `json:"failed_list"`
}
