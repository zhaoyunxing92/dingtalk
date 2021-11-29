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

type GetDriveSpacesFiles struct {
	Response

	//下一页的游标，为空字符串则表示分页结束
	Token string `json:"nextToken"`

	//文件列表
	SpacesFiles []spacesFile `json:"files"`
}

type spacesFile struct {
	//钉盘空间ID
	SpaceId string `json:"spaceId"`

	//父目录ID
	ParentId string `json:"parentId"`

	//文件ID
	FileId string `json:"fileId"`

	//文件名称
	FileName string `json:"fileName"`

	//文件路径
	FilePath string `json:"filePath"`

	//文件类型。
	//file：文件
	//folder：文件夹
	FileType string `json:"fileType"`

	//文件内容类型，取值：
	//
	//image：图片
	//document：一般文档
	//alidoc：阿里文档
	//text：文本
	//video：视频
	//audio：音频
	//archive：归档
	//app：应用
	//link：快捷方式
	//other：其他
	ContentType string `json:"contentType"`

	//文件后缀名
	FileExtension string `json:"fileExtension"`

	//文件大小
	FileSize int `json:"fileSize"`

	//创建时间
	CreateTime string `json:"createTime"`

	//修改时间
	ModifyTime string `json:"modifyTime"`

	//创建者ID
	Creator string `json:"creator"`

	//修改者ID
	Modifier string `json:"modifier"`
}
