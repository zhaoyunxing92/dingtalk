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

	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/file"
)

func TestNewUploadFile(t *testing.T) {
	f := NewUploadFile("../image/dingtalk.png", file.Image)

	assert.NotNil(t, f)
	assert.Equal(t, f.Genre, string(file.Image))
	assert.Equal(t, f.FileName, "dingtalk.png")
	assert.Equal(t, f.FieldName, "media")
}
