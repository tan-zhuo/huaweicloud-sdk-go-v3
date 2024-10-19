package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCustomizedFieldsResponse Response Object
type SearchCustomizedFieldsResponse struct {
	Data           *SearchCustomizedFieldsResultData `json:"data,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o SearchCustomizedFieldsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCustomizedFieldsResponse struct{}"
	}

	return strings.Join([]string{"SearchCustomizedFieldsResponse", string(data)}, " ")
}
