package crypto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSignature(t *testing.T) {
	timestamp := "1629879099342"
	ticket := "SEx31XX20YLqvcLVpnwObQHIbvv85ZD6UGXCphXEmp0zCNNfj6VMMklvZjVNDCf99rrKc476XMKS56d5p6iCOU"
	secret := "c9RMIAGoL_K26gxkw4LrLGf2ERSQVqMBd1K_jclUIT3WM249uVCwf0Ug2RpG4dBh"
	signature := GetSignature(timestamp, secret, ticket)
	sign := "auGsFduirwTx3fYqswR8SCuaKbHpHjJzUB%2FcIAWHGuw%3D"
	assert.Equal(t, sign, signature)
	fmt.Println(signature)
}
