package request

type UnionIdGetUserId struct {
	UnionId string `json:"unionid" validate:"required"`
}

func NewUnionIdGetUserId(unionId string) *UnionIdGetUserId {

	return &UnionIdGetUserId{unionId}
}
