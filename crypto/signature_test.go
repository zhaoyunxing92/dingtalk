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

package crypto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSignature(t *testing.T) {
	timestamp := "1636725918388"
	ticket := "ticket"
	secret := "secret"
	signature := GetSignature(timestamp, secret, ticket)
	sign := "LOZmXJzlQYpIu5emP6/1OF4hE3qlMzY/6y4FbDjXPbU="
	assert.Equal(t, sign, signature)
	fmt.Println(signature)
}
