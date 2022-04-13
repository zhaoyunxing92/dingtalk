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
