package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerResponse Response Object
type UpdateServerResponse struct {
	Server         *UpdateServerResult `json:"server,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerResponse struct{}"
	}

	return strings.Join([]string{"UpdateServerResponse", string(data)}, " ")
}
