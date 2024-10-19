package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStopTransferTaskResponse Response Object
type BatchStopTransferTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchStopTransferTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopTransferTaskResponse struct{}"
	}

	return strings.Join([]string{"BatchStopTransferTaskResponse", string(data)}, " ")
}
