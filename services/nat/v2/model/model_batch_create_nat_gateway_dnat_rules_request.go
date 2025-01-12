package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateNatGatewayDnatRulesRequest Request Object
type BatchCreateNatGatewayDnatRulesRequest struct {
	Body *BatchCreateNatGatewayDnatRulesRequestBody `json:"body,omitempty"`
}

func (o BatchCreateNatGatewayDnatRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateNatGatewayDnatRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateNatGatewayDnatRulesRequest", string(data)}, " ")
}
