package crypto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSignature(t *testing.T) {
	timestamp := "1629879099342"
	ticket := "ticket"
	secret := "secret"
	signature := GetSignature(timestamp, secret, ticket)
	sign := "dvTTVIsYbKXSNcldIxZ1i9Aqx1WATXd2Mnb%2BAe%2Fz07A%3D"
	assert.Equal(t, sign, signature)
	fmt.Println(signature)
}
