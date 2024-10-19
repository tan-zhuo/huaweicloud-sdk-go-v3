package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricRelationsResponse Response Object
type ListMetricRelationsResponse struct {
	Data           *ListMetricRelationsResultData `json:"data,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListMetricRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricRelationsResponse struct{}"
	}

	return strings.Join([]string{"ListMetricRelationsResponse", string(data)}, " ")
}
