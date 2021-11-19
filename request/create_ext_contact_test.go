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
	"github.com/stretchr/testify/assert"
)

func TestNewExtContact(t *testing.T) {

	ec := NewCreateExtContact("张三", "12345678941", "86", "abcde",
		123, 123, 456).
		SetTitle("资深开发").
		SetAddress("杭州").
		SetShareDept(123, 123, 678).
		SetRemark("技术人员").
		SetCompanyName("小番茄").
		SetShareUser("123", "123", "456").
		Build()

	assert.NotEmpty(t, ec)
	assert.Equal(t, len(ec.Contact.ShareUser), 2)
	assert.Equal(t, len(ec.Contact.Labels), 2)
	assert.Equal(t, len(ec.Contact.ShareDept), 2)

	t.Log(ec.String())
}
