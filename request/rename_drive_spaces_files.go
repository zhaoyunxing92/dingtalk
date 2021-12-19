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

type RenameDriveSpacesFiles struct {
	// 用户unionId，可以调用通过免登码获取用户信息(v2)接口获取
	UnionId string `json:"unionId,omitempty" validate:"required"`

	// 新的文件（夹）名称
	NewFileName string `json:"newFileName" validate:"required"`
}

func NewRenameDriveSpacesFiles(newFileName, unionId string) *RenameDriveSpacesFiles {
	return &RenameDriveSpacesFiles{unionId, newFileName}
}
