package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBlockchainNodesRequest Request Object
type ShowBlockchainNodesRequest struct {

	// blockchainID
	BlockchainId string `json:"blockchain_id"`
}

func (o ShowBlockchainNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlockchainNodesRequest struct{}"
	}

	return strings.Join([]string{"ShowBlockchainNodesRequest", string(data)}, " ")
}
