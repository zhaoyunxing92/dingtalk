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
	"errors"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestNewRoleAddUser(t *testing.T) {
	str := NewRoleAddUser([]int{1222, 1111, 1111, 1111}, []string{"123456"})
	err := validate(str)
	if err != nil {
		t.Error(err)
	}
}

// validate 参数验证
func validate(s interface{}) error {
	v := validator.New()
	if err := v.Struct(s); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Error())
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}
