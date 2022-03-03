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
	"github.com/zhaoyunxing92/dingtalk/v2/constant/file"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/policy"
)

type CreateDriveSpacesFiles struct {
	// 钉盘空间ID
	SpaceId string `json:"-" validate:"required"`

	// 用户unionId
	UnionId string `json:"unionId" validate:"required"`

	// 父目录ID,根目录时传0
	ParentId string `json:"parentId,omitempty"`

	//文件类型。
	//
	//file：文件
	//folder：文件夹
	FileType string `json:"fileType,omitempty" validate:"required,oneof=file folder"`

	// 文件名
	FileName string `json:"fileName,omitempty" validate:"required"`

	// 对应OSS Object key。可以通过调用获取文件上传信息接口获取。
	// 当fileType值为file文件时，此参数必填。
	// 当fileType值为folder文件夹时，此参数可不填。
	MediaId string `json:"mediaId,omitempty"`

	//文件名称冲突策略，取值：
	//
	//autoRename：自动重命名
	//overwrite：覆写
	//returnExisting：返回已存在文件
	//returnError：报错
	ConflictPolicy string `json:"addConflictPolicy,omitempty" validate:"omitempty,oneof=autoRename overwrite returnExisting returnError"`
}

type createDriveSpacesFilesBuilder struct {
	file *CreateDriveSpacesFiles
}

func NewCreateDriveSpacesFiles(unionId, spaceId, fileName string, fileType file.Genre) *createDriveSpacesFilesBuilder {
	return &createDriveSpacesFilesBuilder{file: &CreateDriveSpacesFiles{
		UnionId: unionId, SpaceId: spaceId,
		FileName: fileName, FileType: string(fileType),
	}}
}

// SetMediaId 对应OSS Object key。可以通过调用获取文件上传信息接口获取。
// 当fileType值为file文件时，此参数必填。
// 当fileType值为folder文件夹时，此参数可不填。
func (b *createDriveSpacesFilesBuilder) SetMediaId(mediaId string) *createDriveSpacesFilesBuilder {
	b.file.MediaId = mediaId
	return b
}

//SetConflictPolicy 文件名称冲突策略，取值：
//
//autoRename：自动重命名
//overwrite：覆写
//returnExisting：返回已存在文件
//returnError：报错
func (b *createDriveSpacesFilesBuilder) SetConflictPolicy(policy policy.ConflictPolicy) *createDriveSpacesFilesBuilder {
	b.file.ConflictPolicy = string(policy)
	return b
}

// SetParentId 父目录ID,根目录时传0
func (b *createDriveSpacesFilesBuilder) SetParentId(parentId string) *createDriveSpacesFilesBuilder {
	b.file.ParentId = parentId
	return b
}

func (b *createDriveSpacesFilesBuilder) Build() *CreateDriveSpacesFiles {
	return b.file
}
