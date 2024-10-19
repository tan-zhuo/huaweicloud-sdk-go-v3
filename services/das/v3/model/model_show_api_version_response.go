package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApiVersionResponse Response Object
type ShowApiVersionResponse struct {
	Version        *ApiVersion `json:"version,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowApiVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowApiVersionResponse", string(data)}, " ")
}
