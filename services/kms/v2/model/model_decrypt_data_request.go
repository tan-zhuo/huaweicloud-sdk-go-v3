package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DecryptDataRequest Request Object
type DecryptDataRequest struct {
	Body *DecryptDataRequestBody `json:"body,omitempty"`
}

func (o DecryptDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDataRequest struct{}"
	}

	return strings.Join([]string{"DecryptDataRequest", string(data)}, " ")
}
