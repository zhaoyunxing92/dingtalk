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

package role

type SpacesRole string

const (
	//Owner 所有者
	Owner = SpacesRole("owner")

	//Admin 管理员
	Admin = SpacesRole("admin")

	//Editor 可编辑
	Editor = SpacesRole("editor")

	//Viewer 可查看/下载
	Viewer = SpacesRole("viewer")

	//OnlyViewer 只读
	OnlyViewer = SpacesRole("only_viewer")
)
