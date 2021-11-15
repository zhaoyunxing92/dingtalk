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

type MessageProgress struct {
	Response
	progress `json:"progress"`
}

type progress struct {
	//取值0~100，表示处理的百分比
	Percent int `json:"progress_in_percent"`

	//任务执行状态：
	//0：未开始
	//1：处理中
	//2：处理完毕
	Status int `json:"status"`
}
