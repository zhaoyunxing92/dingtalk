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
	"strings"
)

type BatchRemoveUserRole struct {
	RoleIds string `json:"roleIds"`

	RS []int `json:"-" validate:"max=20"`

	UserIds string `json:"userIds"`

	US []string `json:"-" validate:"max=100"`
}

func NewBatchRemoveUserRole(roleIds []int, userIds []string) *BatchRemoveUserRole {
	ds := strings.Join(removeIntDuplicatesToString(roleIds), ",")
	us := strings.Join(removeStringDuplicates(userIds), ",")
	return &BatchRemoveUserRole{
		RoleIds: ds, RS: removeIntDuplicates(roleIds),
		UserIds: us, US: removeStringDuplicates(userIds),
	}
}
