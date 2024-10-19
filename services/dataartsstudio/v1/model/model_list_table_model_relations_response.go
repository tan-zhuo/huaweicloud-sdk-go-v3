package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableModelRelationsResponse Response Object
type ListTableModelRelationsResponse struct {
	Data           *ListTableModelRelationsResultData `json:"data,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListTableModelRelationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableModelRelationsResponse struct{}"
	}

	return strings.Join([]string{"ListTableModelRelationsResponse", string(data)}, " ")
}
