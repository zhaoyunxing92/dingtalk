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

// CorpDingIndex 获取企业最新钉钉指数信息
type CorpDingIndex struct {
	Response
	// 日期
	Date string `json:"statDate"`

	// 钉钉指数
	Total float32 `json:"idxTotal"`

	// 效率指数
	Efficiency float32 `json:"idxEfficiency"`

	// 绿色指数
	Carbon float32 `json:"idxCarbon"`

	// 钉钉指数月均分
	MonthlyAvg float32 `json:"idxMonthlyAvg"`
}
