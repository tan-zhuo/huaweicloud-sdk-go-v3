package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerResponse Response Object
type ShowServerResponse struct {
	Server         *ServerDetail `json:"server,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerResponse struct{}"
	}

	return strings.Join([]string{"ShowServerResponse", string(data)}, " ")
}
