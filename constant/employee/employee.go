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

package employee

//Status 在职员工子状态筛选，可以查询多个状态。不同状态之间使用英文逗号分隔。
//
//2：试用期
//
//3：正式
//
//5：待离职
//
//-1：无状态
type Status string

const (
	// ProbationPeriod 试用期
	ProbationPeriod = Status("2")

	// Formal 正式
	Formal = Status("3")

	// ToBeResigned 待离职
	ToBeResigned = Status("5")

	// Unknown 无状态
	Unknown = Status("-1")
)

//Mold 员工类型枚举值：
//
//0：无类型
//
//1：全职
//
//2：兼职
//
//3：实习
//
//4：劳务派遣
//
//5：退休返聘
//
//6：劳务外包
type Mold string

const (

	// UnknownMold 无类型
	UnknownMold = Mold("0")

	// FullTimeMold 全职
	FullTimeMold = Mold("1")

	// PartTimeMold 兼职
	PartTimeMold = Mold("2")

	// InternshipMold 实习
	InternshipMold = Mold("3")

	// LaborDispatchMold 劳务派遣
	LaborDispatchMold = Mold("4")

	// RescindedMold 退休返聘
	RescindedMold = Mold("5")

	// OutsourcingMold 劳务外包
	OutsourcingMold = Mold("6")
)
