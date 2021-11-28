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
	"github.com/zhaoyunxing92/dingtalk/v2/constant/order"
)

type GetDriveSpacesFiles struct {
	//钉盘空间ID
	SpaceId string `json:"spaceId" validate:"required"`

	//用户unionId
	UnionId string `json:"unionId" validate:"required"`

	//父目录ID
	ParentId string `json:"parentId,omitempty"`

	//分页游标
	Token string `json:"nextToken,omitempty"`

	//分页大小
	Size int `json:"maxResults"`

	//排序类型，取值：
	//
	//createTimeAsc：按创建时间升序
	//createTimeDesc：按创建时间降序
	//modifyTimeAsc：按修改时间升序
	//modifyTimeDesc：按修改时间降序
	//nameAsc：按名称升序
	//nameDesc：按名称降序
	//sizeAsc：按大小升序
	//sizeDesc：按大小降序
	//默认值：createTimeDesc
	OrderType string `json:"orderType,omitempty"`
}

type getDriveSpacesFilesBuilder struct {
	file *GetDriveSpacesFiles
}

func NewGetDriveSpacesFiles(spaceId, unionId string, size int) *getDriveSpacesFilesBuilder {
	return &getDriveSpacesFilesBuilder{file: &GetDriveSpacesFiles{SpaceId: spaceId, UnionId: unionId, Size: size}}
}

func (b *getDriveSpacesFilesBuilder) SetOrderType(order order.OrderType) *getDriveSpacesFilesBuilder {
	b.file.OrderType = string(order)
	return b
}

func (b *getDriveSpacesFilesBuilder) SetToken(token string) *getDriveSpacesFilesBuilder {
	b.file.Token = token
	return b
}

func (b *getDriveSpacesFilesBuilder) SetParentId(parentId string) *getDriveSpacesFilesBuilder {
	b.file.ParentId = parentId
	return b
}

func (b *getDriveSpacesFilesBuilder) Build() *GetDriveSpacesFiles {
	return b.file
}
