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

package message

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestNewMarkDownMessage(t *testing.T) {
	md := NewMarkDownMessage("title", "markdown")
	assert.NotNil(t, md)
	assert.Equal(t, md.Title, "title")
	assert.Equal(t, md.Text, "markdown")
	assert.Equal(t, md.MessageType(), "markdown")
}
