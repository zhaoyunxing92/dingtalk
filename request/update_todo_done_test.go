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

func TestNewUpdateTodoDone(t *testing.T) {
	todo := NewUpdateTodoDone("ABNiSWeAolg5OETyYT60wdQiEiE", "taskc782d9b8cee1127884f04f21ae63e243").
		SetTodoDone("ABNiSWeAolg5OETyYT60wdQiEiE", true).
		SetTodoDone("ABNiSWeAolg5OETyYT60wdQiEiE", false).
		Build()

	assert.NotNil(t, todo)
}
