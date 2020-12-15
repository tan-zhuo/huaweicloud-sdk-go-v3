/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type KeystoneValidateTokenResponse struct {
	Token          *TokenResult `json:"token,omitempty"`
	XSubjectToken  *string      `json:"X-Subject-Token,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o KeystoneValidateTokenResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneValidateTokenResponse", string(data)}, " ")
}