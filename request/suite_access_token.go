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

//SuiteAccessToken 获取suite套件token
type SuiteAccessToken struct {
	Key string `json:"suite_key" validate:"required"`

	Secret string `json:"suite_secret" validate:"required"`

	Ticket string `json:"suite_ticket" validate:"required"`
}

type suiteAccessTokenBuilder struct {
	sat *SuiteAccessToken
}

func NewSuiteAccessToken() *suiteAccessTokenBuilder {
	return &suiteAccessTokenBuilder{sat: &SuiteAccessToken{}}
}

func (s *suiteAccessTokenBuilder) SetKey(key string) *suiteAccessTokenBuilder {
	s.sat.Key = key
	return s
}

func (s *suiteAccessTokenBuilder) SetSecret(secret string) *suiteAccessTokenBuilder {
	s.sat.Secret = secret
	return s
}

func (s *suiteAccessTokenBuilder) SetTicket(ticket string) *suiteAccessTokenBuilder {
	s.sat.Ticket = ticket
	return s
}

func (s *suiteAccessTokenBuilder) Build() *SuiteAccessToken {
	return s.sat
}
