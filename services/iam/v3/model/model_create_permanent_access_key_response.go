package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePermanentAccessKeyResponse Response Object
type CreatePermanentAccessKeyResponse struct {
	Credential     *CreateCredentialResult `json:"credential,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o CreatePermanentAccessKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePermanentAccessKeyResponse struct{}"
	}

	return strings.Join([]string{"CreatePermanentAccessKeyResponse", string(data)}, " ")
}
