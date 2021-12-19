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

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant/file"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/policy"
)

func TestNewCreateDriveSpacesFiles(t *testing.T) {
	f := NewCreateDriveSpacesFiles("unionId", "spaceId", "golang", file.Folder).
		SetParentId("parentId").
		SetConflictPolicy(policy.Overwrite).
		SetMediaId("mediaId").
		Build()

	assert.NotNil(t, f)
	assert.Equal(t, f.MediaId, "mediaId")
	assert.Equal(t, f.SpaceId, "spaceId")
	assert.Equal(t, f.UnionId, "unionId")
	assert.Equal(t, f.FileName, "golang")
	assert.Equal(t, f.ParentId, "parentId")
	assert.Equal(t, f.FileType, string(file.Folder))
	assert.Equal(t, f.ConflictPolicy, string(policy.Overwrite))
}
