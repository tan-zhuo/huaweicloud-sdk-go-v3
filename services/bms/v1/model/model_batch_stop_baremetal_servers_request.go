package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStopBaremetalServersRequest Request Object
type BatchStopBaremetalServersRequest struct {
	Body *OsStopBody `json:"body,omitempty"`
}

func (o BatchStopBaremetalServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopBaremetalServersRequest struct{}"
	}

	return strings.Join([]string{"BatchStopBaremetalServersRequest", string(data)}, " ")
}
