package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProtectedInstanceResponse Response Object
type ShowProtectedInstanceResponse struct {
	ProtectedInstance *ShowProtectedInstanceParams `json:"protected_instance,omitempty"`
	HttpStatusCode    int                          `json:"-"`
}

func (o ShowProtectedInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProtectedInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowProtectedInstanceResponse", string(data)}, " ")
}
