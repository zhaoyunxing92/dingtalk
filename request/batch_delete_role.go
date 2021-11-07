package request

import "strings"

type BatchRemoveUserRole struct {
	RoleIds string `json:"roleIds"`

	RS []int `json:"-" validate:"max=20"`

	UserIds string `json:"userIds"`

	US []string `json:"-" validate:"max=100"`
}

func NewBatchRemoveUserRole(roleIds []int, userIds []string) *BatchRemoveUserRole {
	ds := strings.Join(removeIntDuplicatesToString(roleIds), ",")
	us := strings.Join(removeStringDuplicates(userIds), ",")
	return &BatchRemoveUserRole{RoleIds: ds, RS: removeIntDuplicates(roleIds),
		UserIds: us, US: removeStringDuplicates(userIds)}
}
