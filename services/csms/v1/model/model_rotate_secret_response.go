package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RotateSecretResponse Response Object
type RotateSecretResponse struct {

	// 凭据的版本号标识符。
	VersionId *string `json:"version_id,omitempty"`

	// 凭据的名称。
	SecretName     *string `json:"secret_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RotateSecretResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RotateSecretResponse struct{}"
	}

	return strings.Join([]string{"RotateSecretResponse", string(data)}, " ")
}
