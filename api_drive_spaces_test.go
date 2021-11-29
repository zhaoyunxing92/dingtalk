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
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/zhaoyunxing92/dingtalk/v2/constant/conflict"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/file"
	"github.com/zhaoyunxing92/dingtalk/v2/constant/spaces"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
)

func TestDingTalk_CreateDriveSpaces(t *testing.T) {
	t.Skip()
	res, err := client.CreateDriveSpaces("golang", "ABNiSWeAolg5OETyYT60wdQiEiE")

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_DeleteDriveSpaces(t *testing.T) {
	t.Skip()
	res, err := client.DeleteDriveSpaces("5296669223", "ABNiSWeAolg5OETyYT60wdQiEiE")

	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetDriveSpaces(t *testing.T) {
	res, err := client.GetDriveSpaces("ABNiSWeAolg5OETyYT60wdQiEiE", spaces.Org, "", 5)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(res.Spaces), 5)
}

func TestDingTalk_GetDriveSpacesInfo(t *testing.T) {
	res, err := client.GetDriveSpacesInfo("3452011774", "ABNiSWeAolg5OETyYT60wdQiEiE")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.Name, "全员文件夹")
}

func TestDingTalk_GetDriveSpacesFiles(t *testing.T) {
	files := request.NewGetDriveSpacesFiles("3452011774", "ABNiSWeAolg5OETyYT60wdQiEiE", 10).
		Build()

	res, err := client.GetDriveSpacesFiles(files)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestDingTalk_GetDriveSpacesFileInfo(t *testing.T) {

	res, err := client.GetDriveSpacesFileInfo("3452011774", "47453824236",
		"ABNiSWeAolg5OETyYT60wdQiEiE")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.FileExtension, "pdf")
}

func TestDingTalk_CreateDriveSpacesFiles(t *testing.T) {

	f := request.NewCreateDriveSpacesFiles("ABNiSWeAolg5OETyYT60wdQiEiE", "3452011774",
		"golang", file.Folder).
		SetConflictPolicy(conflict.Overwrite).
		Build()

	res, err := client.CreateDriveSpacesFiles(f)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, res.FileName, "golang")
}
