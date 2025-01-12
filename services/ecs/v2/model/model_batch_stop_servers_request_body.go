package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStopServersRequestBody This is a auto create Body Object
type BatchStopServersRequestBody struct {
	OsStop *BatchStopServersOption `json:"os-stop"`
}

func (o BatchStopServersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopServersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchStopServersRequestBody", string(data)}, " ")
}
