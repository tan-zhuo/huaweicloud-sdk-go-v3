package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClickHouseInstanceResponse Response Object
type ListClickHouseInstanceResponse struct {
	Instance       *ChInstancesInfoRsponseInstance `json:"instance,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o ListClickHouseInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClickHouseInstanceResponse struct{}"
	}

	return strings.Join([]string{"ListClickHouseInstanceResponse", string(data)}, " ")
}
