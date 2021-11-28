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

package dingtalk

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/spaces"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"github.com/zhaoyunxing92/dingtalk/v2/response"
)

//CreateDriveSpaces 新建空间
func (ding *dingTalk) CreateDriveSpaces(name, unionId string) (rsp response.CreateDriveSpaces, err error) {

	return rsp, ding.Request(http.MethodPost, constant.CreateDriveSpacesKey, nil,
		request.NewCreateDriveSpaces(name, unionId), &rsp)
}

//DeleteDriveSpaces 删除空间
func (ding *dingTalk) DeleteDriveSpaces(spaceId, unionId string) (rsp response.DeleteDriveSpaces, err error) {

	query := url.Values{}
	query.Set("unionId", unionId)

	return rsp, ding.Request(http.MethodDelete, fmt.Sprintf(constant.DeleteDriveSpacesKey, spaceId), query, nil, &rsp)
}

//GetDriveSpaces 获取空间列表
func (ding *dingTalk) GetDriveSpaces(unionId string, spaceType spaces.SpaceType, token string, size int) (rsp response.GetDriveSpaces, err error) {

	query := url.Values{}
	query.Set("spaceType", string(spaceType))
	query.Set("unionId", unionId)
	query.Set("nextToken", token)
	query.Set("maxResults", strconv.Itoa(size))

	return rsp, ding.Request(http.MethodGet, constant.GetDriveSpacesKey, query, nil, &rsp)
}

//GetDriveSpacesInfo 获取空间信息
func (ding *dingTalk) GetDriveSpacesInfo(spaceId, unionId string) (rsp response.CreateDriveSpaces, err error) {

	query := url.Values{}
	query.Set("unionId", unionId)

	return rsp, ding.Request(http.MethodGet, fmt.Sprintf(constant.GetDriveSpacesInfoKey, spaceId), query, nil, &rsp)
}

//GetDriveSpacesFiles 查询文件（夹）列表
//order 取值：
//createTimeAsc：按创建时间升序
//createTimeDesc：按创建时间降序
//modifyTimeAsc：按修改时间升序
//modifyTimeDesc：按修改时间降序
//nameAsc：按名称升序
//nameDesc：按名称降序
//sizeAsc：按大小升序
//sizeDesc：按大小降序
//默认值：createTimeDesc
func (ding *dingTalk) GetDriveSpacesFiles(res *request.GetDriveSpacesFiles) (rsp response.GetDriveSpacesFiles, err error) {

	query := url.Values{}
	query.Set("parentId", res.ParentId)
	query.Set("unionId", res.UnionId)
	query.Set("nextToken", res.Token)
	query.Set("maxResults", strconv.Itoa(res.Size))
	query.Set("orderType", res.OrderType)

	return rsp, ding.Request(http.MethodGet, fmt.Sprintf(constant.GetDriveSpacesFilesKey, res.SpaceId), query,
		nil, &rsp)
}

//GetDriveSpacesFileInfo 查询文件（夹）信息
func (ding *dingTalk) GetDriveSpacesFileInfo(spaceId, fileId, unionId string) (rsp response.GetDriveSpacesFileInfo, err error) {
	query := url.Values{}
	query.Set("unionId", unionId)

	return rsp, ding.Request(http.MethodGet, fmt.Sprintf(constant.GetDriveSpacesFileInfoKey, spaceId, fileId),
		query, nil, &rsp)
}
