package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApiVersionRequest Request Object
type ShowApiVersionRequest struct {

	// API版本
	Version string `json:"version"`
}

func (o ShowApiVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowApiVersionRequest", string(data)}, " ")
}
