package request

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
	"testing"
)

func TestNewRoleAddUser(t *testing.T) {

	str := NewRoleAddUser([]int{1222, 1111, 1111, 1111}, []string{"123456"})
	err := validate(str)
	if err != nil {
		t.Error(err)
	}
}

// validate 参数验证
func validate(s interface{}) error {
	v := validator.New()
	if err := v.Struct(s); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Error())
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}
