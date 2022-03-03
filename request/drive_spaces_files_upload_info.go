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

type GetDriveSpacesFilesUploadInfo struct {
	// 钉盘空间ID
	SpaceId string `json:"-" validate:"required"`

	// ParentId 父目录ID。
	ParentId string `json:"-" validate:"required"`

	UnionId string `json:"unionId,omitempty" validate:"required"`

	// 文件名，带文件扩展名
	FileName string `json:"fileName,omitempty" validate:"required"`

	// 文件大小
	FileSize int `json:"fileSize,omitempty" validate:"required"`

	// 文件md5
	Md5 string `json:"md5,omitempty" validate:"required"`

	//文件名称冲突策略，取值：
	//
	//autoRename：自动重命名
	//overwrite：覆写
	//returnExisting：返回已存在文件
	//returnError：报错
	ConflictPolicy string `json:"addConflictPolicy,omitempty" validate:"omitempty,oneof=autoRename overwrite returnExisting returnError"`

	//对应OSS Object key。
	//
	//如果首次调用本接口，为了获取文件的上传信息，无需填写mediaId。
	//如果非首次调用本接口，为了刷新上次调用时返回值中的accessToken，需要传入上次调用返回值的mediaId。
	MediaId string `json:"mediaId,omitempty"`
}

type driveSpacesFilesUploadInfoBuilder struct {
	info *GetDriveSpacesFilesUploadInfo
}

func NewGetDriveSpacesFilesUploadInfo(spaceId, parenId, fileName string, fileSize int, md5,
	unionId string,
) *driveSpacesFilesUploadInfoBuilder {
	return &driveSpacesFilesUploadInfoBuilder{info: &GetDriveSpacesFilesUploadInfo{
		SpaceId: spaceId, ParentId: parenId,
		FileName: fileName, FileSize: fileSize, Md5: md5, UnionId: unionId,
	}}
}
