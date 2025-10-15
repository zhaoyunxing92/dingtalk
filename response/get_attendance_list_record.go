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

type GetAttendanceListRecord struct {
	Response

	AttendanceListRecord []struct {
		// 用户打卡定位精度
		UserAccuracy float32 `json:"userAccuracy"`

		// 班次ID
		ClassId int64 `json:"classId"`

		// 用户打卡纬度说明: 打卡数据来源为ATM或DING_ATM，不返回该字段
		UserLatitude float32 `json:"userLatitude"`

		// 用户打卡经度说明: 打卡数据来源为ATM或DING_ATM，不返回该字段
		UserLongitude float32 `json:"userLongitude"`

		// 用户打卡地址说明: 如果是考勤机打卡 userAddress，返回的是考勤机名称
		UserAddress string `json:"userAddress"`

		// 打卡设备ID说明: 使用手机钉钉客户端打卡返回该字段
		DeviceId string `json:"deviceId"`

		// 定位方法
		LocationMethod string `json:"locationMethod"`

		// 是否合法
		// - Y：合法， 说明: 当timeResult和locationResult都为Normal时，为该值
		// - N：不合法
		IsLegal string `json:"isLegal"`

		// 实际打卡时间
		UserCheckTime int64 `json:"userCheckTime"`

		// 关联的审批实例ID，当该字段非空时，表示打卡记录与请假、加班等审批有关
		ProcInstId string `json:"procInstId"`

		// 计算迟到和早退，基准时间；也可作为排班打卡时间
		BaseCheckTime int64 `json:"baseCheckTime"`

		// 关联的审批ID，当该字段非空时，表示打卡记录与请假、加班等审批有关
		ApproveId int64 `json:"approveId"`

		// 打卡结果：
		// Normal：正常
		// Early：早退
		// Late：迟到
		// SeriousLate：严重迟到
		// Absenteeism：旷工迟到
		// NotSigned：未打卡
		TimeResult string `json:"timeResult"`

		// 位置结果：
		// - Normal：范围内
		// - Outside：范围外
		// - NotSigned：未打卡
		LocationResult string `json:"locationResult"`

		// 考勤类型：
		// - OnDuty：上班
		// - OffDuty：下班
		CheckType string `json:"checkType"`

		// 数据来源：
		// - ATM：考勤机打卡（指纹/人脸打卡）
		// - BEACON：IBeacon
		// - DING_ATM：钉钉考勤机（考勤机蓝牙打卡）
		// - USER：用户打卡
		// - BOSS：老板改签
		// - APPROVE：审批系统
		// - SYSTEM：考勤系统
		// - AUTO_CHECK：自动打卡
		SourceType string `json:"sourceType"`

		// 打卡人的userId
		UserId string `json:"userId"`

		// 工作日
		WorkDate int64 `json:"workDate"`

		// 企业ID
		CorpId string `json:"corpId"`

		// 排班ID
		PlanId int64 `json:"planId"`

		// 考勤组ID
		GroupId int64 `json:"groupId"`

		// 考勤ID
		Id int64 `json:"id"`

		// 异常信息类型：
		// - Security：安全相关原因
		// - Other：其他原因
		InvalidRecordType string `json:"invalidRecordType"`

		// 用户打卡wifi SSID
		UserSsid string `json:"userSsid"`

		// 用户打卡wifi Mac地址
		UserMacAddr string `json:"userMacAddr"`

		// 排班打卡时间
		PlanCheckTime int64 `json:"planCheckTime"`

		// 基准地址
		BaseAddress string `json:"baseAddress"`

		// 基准经度说明: 打卡数据来源为ATM或DING_ATM，不返回该字段
		BaseLongitude string `json:"baseLongitude"`

		// 基准纬度说明: 打卡数据来源为ATM或DING_ATM，不返回该字段
		BaseLatitude string `json:"baseLatitude"`

		// 基准定位精度
		BaseAccuracy string `json:"baseAccuracy"`

		// 基准wifi ssid
		BaseSsid string `json:"baseSsid"`

		// 基准Mac地址
		BaseMacAddr string `json:"baseMacAddr"`

		// 打卡记录创建时间
		GmtCreate int64 `json:"gmtCreate"`

		// 对应的invalidRecordType异常信息的具体描述
		InvalidRecordMsg string `json:"invalidRecordMsg"`

		// 打卡记录修改时间
		GmtModified int64 `json:"gmtModified"`

		// 打卡备注
		OutsideRemark string `json:"outsideRemark"`

		// 打卡设备序列号
		DeviceSN string `json:"deviceSN"`

		// 关联的业务ID
		BizId string `json:"bizId"`

		// 拍照图片 URL说明: 打卡详情中有拍照的图片信息，会返回 photoUrl，无拍照图片信息，则该字段不返回
		PhotoUrl string `json:"photoUrl"`
	} `json:"recordresult"`
}
