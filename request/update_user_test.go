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

package request

import (
	"testing"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant/language"
)

func TestNewUpdateUser(t *testing.T) {

	str := NewUpdateUser("134567").
		SetForceUpdateFields("").
		SetHideMobile(false).
		SetSeniorMode(false).
		SetHiredDate(1).
		SetLanguage(language.EN_US).
		SetExtension("ext").
		SetLoginId("").
		SetDeptTitle(123, "开发").
		SetDept(123, 456, 789).
		SetWorkPlace("work").
		SetRemark("remark").
		SetOrgEmailType("remark").
		SetOrgEmail("remark").
		SetEmail("remark").
		SetTitle("remark").
		SetMobile("remark").
		SetName("remark").
		SetManagerUserId("remark").
		SetTelephone("remark").
		SetJobNumber("remark").
		SetDeptOrder(123, 1).
		SetForceUpdateFields("").
		Build().
		String()

	t.Log(str)
}
