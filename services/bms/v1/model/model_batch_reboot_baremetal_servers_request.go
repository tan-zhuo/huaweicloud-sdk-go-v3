package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRebootBaremetalServersRequest Request Object
type BatchRebootBaremetalServersRequest struct {
	Body *RebootBody `json:"body,omitempty"`
}

func (o BatchRebootBaremetalServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootBaremetalServersRequest struct{}"
	}

	return strings.Join([]string{"BatchRebootBaremetalServersRequest", string(data)}, " ")
}
