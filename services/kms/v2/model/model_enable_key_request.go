package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableKeyRequest Request Object
type EnableKeyRequest struct {
	Body *OperateKeyRequestBody `json:"body,omitempty"`
}

func (o EnableKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableKeyRequest struct{}"
	}

	return strings.Join([]string{"EnableKeyRequest", string(data)}, " ")
}
