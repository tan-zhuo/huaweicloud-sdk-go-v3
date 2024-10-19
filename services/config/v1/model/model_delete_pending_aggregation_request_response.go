package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePendingAggregationRequestResponse Response Object
type DeletePendingAggregationRequestResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePendingAggregationRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePendingAggregationRequestResponse struct{}"
	}

	return strings.Join([]string{"DeletePendingAggregationRequestResponse", string(data)}, " ")
}
