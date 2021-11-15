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

//CreateDriveSpaces  新建空间
type CreateDriveSpaces struct {
	//空间名称（不能为空）
	Name string `json:"name,omitempty" validate:"required"`

	//用户unionId
	UnionId string `json:"unionId,omitempty" validate:"required"`
}

func NewCreateDriveSpaces(name, unionId string) *CreateDriveSpaces {
	return &CreateDriveSpaces{name, unionId}
}
