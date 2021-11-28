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

type OrderType string

const (
	// CreateTimeAsc 按创建时间升序
	CreateTimeAsc = OrderType("createTimeAsc")

	// CreateTimeDesc 按创建时间降序
	CreateTimeDesc = OrderType("createTimeDesc")

	//ModifyTimeAsc 按修改时间升序
	ModifyTimeAsc = OrderType("modifyTimeAsc")

	//ModifyTimeDesc 按修改时间降序
	ModifyTimeDesc = OrderType("modifyTimeDesc")

	//NameAsc 按名称升序
	NameAsc = OrderType("nameAsc")

	//NameDesc 按名称降序
	NameDesc = OrderType("nameDesc")

	//SizeAsc 按大小升序
	SizeAsc = OrderType("sizeAsc")

	//SizeDesc 按大小降序
	SizeDesc = OrderType("sizeDesc")
)
