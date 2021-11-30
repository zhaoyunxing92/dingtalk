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
	"github.com/zhaoyunxing92/dingtalk/v2/constant/policy"
)

func TestNewMoveDriveSpacesFiles(t *testing.T) {
	files := NewMoveDriveSpacesFiles("spaceId", "targetSpaceId", "fileId", "unionId").
		SetTargetParentId("targetParentId").
		SetConflictPolicy(policy.Overwrite).
		Build()

	assert.NotNil(t, files)
	assert.Equal(t, files.FileId, "fileId")
	assert.Equal(t, files.FileId, "fileId")
	assert.Equal(t, files.SpaceId, "spaceId")
	assert.Equal(t, files.UnionId, "unionId")
	assert.Equal(t, files.TargetSpaceId, "targetSpaceId")
	assert.Equal(t, files.TargetParentId, "targetParentId")
	assert.Equal(t, files.ConflictPolicy, string(policy.Overwrite))
}
