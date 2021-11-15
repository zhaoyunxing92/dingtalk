/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAdminUserScope(t *testing.T) {

	scope := NewAdminUserScope("123456")

	assert.NotNil(t, scope)
}

func TestAdminUserScope_String(t *testing.T) {
	scope := NewAdminUserScope("123456").String()

	assert.NotNil(t, scope)
}
