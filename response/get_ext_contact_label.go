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

type GetExtContactLabel struct {
	Response

	Results []labels `json:"results"`
}

type labels struct {
	// 标签组名字
	Name string `json:"name"`

	// 标签组颜色
	Color int `json:"color"`

	// 标签
	Labels []label `json:"labels"`
}

type label struct {
	// 标签名称
	Name string `json:"name"`

	// 标签id
	Id int `json:"id"`
}
