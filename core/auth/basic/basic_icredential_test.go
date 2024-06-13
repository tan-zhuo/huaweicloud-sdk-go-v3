// Copyright 2022 Huawei Technologies Co.,Ltd.
//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package basic

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCredentials_NeedUpdate(t *testing.T) {
	credentials, err := NewCredentialsBuilder().SafeBuild()
	assert.Nil(t, err)
	// Manually specifying a security token
	credentials, err = NewCredentialsBuilder().
		WithAk("ak").
		WithSk("sk").
		WithSecurityToken("token").
		SafeBuild()
	assert.Nil(t, err)
	assert.False(t, credentials.needUpdateSecurityToken())
	// Automatically update the expired security token
	credentials.expiredAt = 1
	assert.True(t, credentials.needUpdateSecurityToken())
	// The security token has not expired
	credentials.expiredAt = time.Now().Unix() + 120
	assert.False(t, credentials.needUpdateSecurityToken())
}

func TestCredentials_NeedUpdateAuthToken(t *testing.T) {
	credentials, err := NewCredentialsBuilder().SafeBuild()
	assert.Nil(t, err)
	// Without idp_id or id_token
	assert.False(t, credentials.needUpdateAuthToken())
	credentials.IdpId = "idp_id"
	credentials.IdTokenFile = "file"
	// Get the auth token for the first time
	assert.True(t, credentials.needUpdateAuthToken())
	// Automatically update the expired auth token
	credentials.authToken = "auth_token"
	assert.True(t, credentials.needUpdateAuthToken())
	// The auth token has not expired
	credentials.expiredAt = time.Now().Unix() + 120
	assert.False(t, credentials.needUpdateAuthToken())
}

// test build with IdpId, IdTokenFile and ProjectId
func TestCredentialsBuilder_SafeBuild(t *testing.T) {
	// ProjectId is missing
	_, err := NewCredentialsBuilder().WithIdpId("id").WithIdTokenFile("file").SafeBuild()
	assert.IsType(t, err, &sdkerr.CredentialsTypeError{})
	assert.Equal(t, "ProjectId is required when using IdpId&IdTokenFile", err.(*sdkerr.CredentialsTypeError).ErrorMessage)
	// IdTokenFile is missing
	_, err = NewCredentialsBuilder().WithIdpId("id").SafeBuild()
	assert.IsType(t, err, &sdkerr.CredentialsTypeError{})
	assert.Equal(t, "IdTokenFile is required when using IdpId&IdTokenFile", err.(*sdkerr.CredentialsTypeError).ErrorMessage)
	// IdpId is missing
	_, err = NewCredentialsBuilder().WithIdTokenFile("file").SafeBuild()
	assert.IsType(t, err, &sdkerr.CredentialsTypeError{})
	assert.Equal(t, "IdpId is required when using IdpId&IdTokenFile", err.(*sdkerr.CredentialsTypeError).ErrorMessage)
	// success with IdpId, IdTokenFile and ProjectId
	credentials, err := NewCredentialsBuilder().WithIdpId("id").WithIdTokenFile("file").WithProjectId("projectId").SafeBuild()
	assert.Nil(t, err)
	assert.Equal(t, "id", credentials.IdpId)
	assert.Equal(t, "file", credentials.IdTokenFile)
	assert.Equal(t, "projectId", credentials.ProjectId)
}

// test empty ak&sk
func TestCredentialsBuilder_SafeBuild2(t *testing.T) {
	// ak is empty string
	_, err := NewCredentialsBuilder().WithAk("").WithSk("sk").SafeBuild()
	assert.IsType(t, err, &sdkerr.CredentialsTypeError{})
	assert.Contains(t, err.(*sdkerr.CredentialsTypeError).ErrorMessage, "input ak cannot be an empty string")
	// sk is empty string
	_, err = NewCredentialsBuilder().WithAk("ak").WithSk("").SafeBuild()
	assert.IsType(t, err, &sdkerr.CredentialsTypeError{})
	assert.Contains(t, err.(*sdkerr.CredentialsTypeError).ErrorMessage, "input sk cannot be an empty string")
	// ak and sk are both empty string
	_, err = NewCredentialsBuilder().WithAk("").WithSk("").SafeBuild()
	assert.IsType(t, err, &sdkerr.CredentialsTypeError{})
	assert.Contains(t, err.(*sdkerr.CredentialsTypeError).ErrorMessage, "input ak cannot be an empty string")
	assert.Contains(t, err.(*sdkerr.CredentialsTypeError).ErrorMessage, "input sk cannot be an empty string")
	// success with valid ak and sk
	credentials, err := NewCredentialsBuilder().WithAk("ak").WithSk("sk").SafeBuild()
	assert.Nil(t, err)
	assert.Equal(t, "ak", credentials.AK)
	assert.Equal(t, "sk", credentials.SK)
	credentials, err = NewCredentialsBuilder().WithAk("").WithSk("").WithAk("ak").WithSk("sk").SafeBuild()
	assert.Nil(t, err)
	assert.Equal(t, "ak", credentials.AK)
	assert.Equal(t, "sk", credentials.SK)
}
