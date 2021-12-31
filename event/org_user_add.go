package event

//OrgUserAdd 通讯录用户增加
//{
//    "CorpId": "dingc7c5220402493357f2c783f7214b6d69",
//    "EventType": "user_add_org",
//    "UserId": [
//        "184919295227658120"
//    ],
//    "OptStaffId": "manager164",
//    "TimeStamp": "1640921671286"
//}
type OrgUserAdd struct {
	OrgUserModify
}
