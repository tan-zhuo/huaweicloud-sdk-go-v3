package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartServersRequest Request Object
type BatchStartServersRequest struct {
	Body *BatchStartServersRequestBody `json:"body,omitempty"`
}

func (o BatchStartServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartServersRequest struct{}"
	}

	return strings.Join([]string{"BatchStartServersRequest", string(data)}, " ")
}
