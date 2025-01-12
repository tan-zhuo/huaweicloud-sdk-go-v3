package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePermanentAccessKeyResponse Response Object
type UpdatePermanentAccessKeyResponse struct {
	Credential     *UpdateCredentialResult `json:"credential,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o UpdatePermanentAccessKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePermanentAccessKeyResponse struct{}"
	}

	return strings.Join([]string{"UpdatePermanentAccessKeyResponse", string(data)}, " ")
}
