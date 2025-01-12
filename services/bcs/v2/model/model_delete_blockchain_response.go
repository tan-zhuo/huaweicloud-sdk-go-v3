package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBlockchainResponse Response Object
type DeleteBlockchainResponse struct {

	// 操作记录id
	OperationId    *string `json:"operation_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteBlockchainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBlockchainResponse struct{}"
	}

	return strings.Join([]string{"DeleteBlockchainResponse", string(data)}, " ")
}
