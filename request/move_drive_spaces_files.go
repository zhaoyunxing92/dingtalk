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
	"github.com/zhaoyunxing92/dingtalk/v2/constant/policy"
)

type MoveDriveSpacesFiles struct {
	// 空间ID。可调用获取空间列表接口获取
	SpaceId string `json:"-" validate:"required"`

	// 文件ID。可以调用查询文件列表接口获取
	FileId string `json:"-" validate:"required"`

	// 用户unionId，可以调用通过免登码获取用户信息(v2)接口获取
	UnionId string `json:"unionId,omitempty" validate:"required"`

	// 目标空间ID
	TargetSpaceId string `json:"targetSpaceId,omitempty" validate:"required"`

	// 目标父目录ID
	TargetParentId string `json:"targetParentId,omitempty"`

	//文件名称冲突策略，取值：
	//
	//autoRename：自动重命名
	//overwrite：覆写
	//returnExisting：返回已存在文件
	//returnError：报错
	ConflictPolicy string `json:"addConflictPolicy,omitempty" validate:"omitempty,oneof=autoRename overwrite returnExisting returnError"`
}

type moveDriveSpacesFilesBuilder struct {
	file *MoveDriveSpacesFiles
}

func NewMoveDriveSpacesFiles(spaceId, fileId, targetSpaceId, unionId string) *moveDriveSpacesFilesBuilder {
	return &moveDriveSpacesFilesBuilder{file: &MoveDriveSpacesFiles{
		SpaceId: spaceId, TargetSpaceId: targetSpaceId,
		FileId: fileId, UnionId: unionId,
	}}
}

//SetConflictPolicy 文件名称冲突策略，取值：
//
//autoRename：自动重命名
//overwrite：覆写
//returnExisting：返回已存在文件
//returnError：报错
func (b *moveDriveSpacesFilesBuilder) SetConflictPolicy(policy policy.ConflictPolicy) *moveDriveSpacesFilesBuilder {
	b.file.ConflictPolicy = string(policy)
	return b
}

func (b *moveDriveSpacesFilesBuilder) SetTargetParentId(targetParentId string) *moveDriveSpacesFilesBuilder {
	b.file.TargetParentId = targetParentId
	return b
}

func (b *moveDriveSpacesFilesBuilder) Build() *MoveDriveSpacesFiles {
	return b.file
}
