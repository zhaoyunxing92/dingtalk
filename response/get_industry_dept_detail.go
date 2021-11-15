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

package response

type GetIndustryDeptDetail struct {
	Response

	industryDeptDetail `json:"result"`
}

type industryDeptDetail struct {
	//部门名称
	Name string `json:"name"`
	//父部门ID
	SuperId int `json:"super_id"`

	//部门类型。行业相关。例如：
	//
	//针对学校
	//
	//campus：校区
	//
	//period：学段
	//
	//grade：年级
	//
	//class：班级
	//
	//针对农村
	//
	//VillageGroup：组
	//
	//Residence：户
	//
	//LeaseholderDept：租客
	//
	//SecretaryDept：村委
	DeptType string `json:"dept_type"`

	//通讯录类型 行业相关。例如：
	//
	//针对学校
	//
	//classic：传统经典校区、学段、年级、班级4层结构。
	//
	//custom：自定义结构
	//
	//针对农村
	//
	//Origin：传统农村类型
	//
	//Community：社区类型
	//
	//custom：自定义结构
	ContactType string `json:"contact_type"`

	//部门的其他业务属性。可JSON反序列化。例如：
	//
	//针对家校
	//
	//period_type：学段类型（幼儿园，小学等）
	//
	//name_mode：学段对应的名称类型（一年级，一年级1班等）
	//
	//grade_level：年纪级数（一年级值为1）
	//
	//start_year：入学年份
	//
	//class_level：年级下班级级数
	//
	//针对农村
	//
	//manager_user_id：组长userID
	//
	//home_tel：家庭电话
	//
	//destitute：是否贫困户
	Feature string `json:"feature"`
}
