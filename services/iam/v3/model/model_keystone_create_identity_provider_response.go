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
type KeystoneCreateIdentityProviderResponse struct {
	IdentityProvider *IdentityprovidersResult `json:"identity_provider,omitempty"`
	HttpStatusCode   int                      `json:"-"`
}

func (o KeystoneCreateIdentityProviderResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneCreateIdentityProviderResponse", string(data)}, " ")
}