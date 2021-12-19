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

package order

// FileOrder 文件排序
type FileOrder string

// DeptOrder 部门排序
type DeptOrder string

const (
	// CreateTimeAsc 按创建时间升序
	CreateTimeAsc = FileOrder("createTimeAsc")

	// CreateTimeDesc 按创建时间降序
	CreateTimeDesc = FileOrder("createTimeDesc")

	// ModifyTimeAsc 按修改时间升序
	ModifyTimeAsc = FileOrder("modifyTimeAsc")

	// ModifyTimeDesc 按修改时间降序
	ModifyTimeDesc = FileOrder("modifyTimeDesc")

	// NameAsc 按名称升序
	NameAsc = FileOrder("nameAsc")

	// NameDesc 按名称降序
	NameDesc = FileOrder("nameDesc")

	// SizeAsc 按大小升序
	SizeAsc = FileOrder("sizeAsc")

	// SizeDesc 按大小降序
	SizeDesc = FileOrder("sizeDesc")
)

const (
	// EntryAsc 代表按照进入部门的时间升序
	EntryAsc = DeptOrder("entry_asc")

	// EntryDesc 代表按照进入部门的时间降序
	EntryDesc = DeptOrder("entry_desc")

	// ModifyAsc 代表按照部门信息修改时间升序
	ModifyAsc = DeptOrder("modify_asc")

	// ModifyDesc 代表按照部门信息修改时间降序
	ModifyDesc = DeptOrder("modify_desc")

	// Custom 代表用户定义(未定义时按照拼音)排序
	Custom = DeptOrder("custom")
)
