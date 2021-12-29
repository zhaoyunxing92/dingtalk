package event

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrmCustomerTrack(t *testing.T) {
	str := `{
    "CorpId": "dingf5d3534a855bdef9a39a90f97fcb1e09",
    "corpId": "dingf5d3534a855bdef9a39a90f97fcb1e09",
    "EventType": "crm_customer_track",
    "EventTime": 1640749233707,
    "BizId": "21057D5806944EF9C93192B36DC5F7B4",
    "tracks": [
        {
            "creator": "manager4268",
            "corpId": "dingf5d3534a855bdef9a39a90f97fcb1e09",
            "customerId": "ad6b2b05-91f1-4393-8ed5-96e9e558b604",
            "subType": 1,
            "id": "qoZCjgT7of7XAtUjo/4JgFy+RR432dJXcXqRbNxwTVM=",
            "gmtCreate": 1640749233707,
            "type": 1
        }
    ]
}`

	crm := &CrmCustomerTrack{}

	err := json.Unmarshal([]byte(str), crm)

	assert.Nil(t, err)
	assert.Equal(t, crm.TimeStamp, 1640749233707)
	assert.Equal(t, crm.EventType, "crm_customer_track")
	assert.Equal(t, crm.CorpId, "dingf5d3534a855bdef9a39a90f97fcb1e09")
	assert.Equal(t, crm.BizId, "21057D5806944EF9C93192B36DC5F7B4")

	assert.Equal(t, len(crm.Tracks), 1)

	assert.Equal(t, crm.Tracks[0].Type, 1)
	assert.Equal(t, crm.Tracks[0].SubType, 1)
	assert.Equal(t, crm.Tracks[0].GmtCreate, 1640749233707)
	assert.Equal(t, crm.Tracks[0].Creator, "manager4268")
	assert.Equal(t, crm.Tracks[0].CorpId, "dingf5d3534a855bdef9a39a90f97fcb1e09")
	assert.Equal(t, crm.Tracks[0].Id, "qoZCjgT7of7XAtUjo/4JgFy+RR432dJXcXqRbNxwTVM=")
	assert.Equal(t, crm.Tracks[0].CustomerId, "ad6b2b05-91f1-4393-8ed5-96e9e558b604")
}
