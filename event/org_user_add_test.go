package event

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrgUserAdd(t *testing.T) {
	str := `{
    "CorpId": "dingc7c5220402493357f2c783f7214b6d69",
    "EventType": "user_add_org",
    "UserId": [
        "184919295227658120"
    ],
    "OptStaffId": "manager164",
    "TimeStamp": "1640921671286"
}`
	user := &OrgUserAdd{}

	err := json.Unmarshal([]byte(str), user)

	assert.Nil(t, err)
	assert.Equal(t, user.EventType, "user_add_org")
	assert.Equal(t, user.OptStaffId, "manager164")
	assert.Equal(t, user.TimeStamp, 1640921671286)
	assert.Equal(t, len(user.UserIds), 1)
	assert.Equal(t, user.CorpId, "dingc7c5220402493357f2c783f7214b6d69")
}
