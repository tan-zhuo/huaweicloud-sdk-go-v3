package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNewBlockchainRequest Request Object
type CreateNewBlockchainRequest struct {
	Body *CreateRequestBody `json:"body,omitempty"`
}

func (o CreateNewBlockchainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNewBlockchainRequest struct{}"
	}

	return strings.Join([]string{"CreateNewBlockchainRequest", string(data)}, " ")
}
