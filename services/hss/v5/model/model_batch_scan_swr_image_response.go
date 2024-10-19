package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchScanSwrImageResponse Response Object
type BatchScanSwrImageResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchScanSwrImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchScanSwrImageResponse struct{}"
	}

	return strings.Join([]string{"BatchScanSwrImageResponse", string(data)}, " ")
}
