package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneCreateProjectResponse Response Object
type KeystoneCreateProjectResponse struct {
	Project        *AuthProjectResult `json:"project,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o KeystoneCreateProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateProjectResponse struct{}"
	}

	return strings.Join([]string{"KeystoneCreateProjectResponse", string(data)}, " ")
}
