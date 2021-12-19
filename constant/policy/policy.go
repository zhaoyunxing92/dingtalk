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

package policy

// ConflictPolicy 文件名称冲突策略
type ConflictPolicy string

// DeletePolicy 删除策略
type DeletePolicy string

// 文件名称冲突策略
const (
	// AutoRename 自动重命名
	AutoRename = ConflictPolicy("autoRename")

	// Overwrite 覆写
	Overwrite = ConflictPolicy("overwrite")

	// ReturnExisting 返回已存在文件
	ReturnExisting = ConflictPolicy("returnExisting")

	// ReturnError 返回已存在文件
	ReturnError = ConflictPolicy("returnError")
)

// 删除策略
const (
	// ToRecycle 删除到回收站
	ToRecycle = DeletePolicy("toRecycle")

	// Completely 彻底删除
	Completely = DeletePolicy("completely")
)
