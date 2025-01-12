package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBcsMetricRequest Request Object
type ListBcsMetricRequest struct {

	// 区块链服务id
	BlockchainId string `json:"blockchain_id"`

	Body *ListBcsMetricRequestBody `json:"body,omitempty"`
}

func (o ListBcsMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBcsMetricRequest struct{}"
	}

	return strings.Join([]string{"ListBcsMetricRequest", string(data)}, " ")
}
