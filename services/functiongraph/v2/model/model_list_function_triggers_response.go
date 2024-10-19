package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFunctionTriggersResponse Response Object
type ListFunctionTriggersResponse struct {
	Body           *[]ListFunctionTriggerResult `json:"body,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListFunctionTriggersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionTriggersResponse struct{}"
	}

	return strings.Join([]string{"ListFunctionTriggersResponse", string(data)}, " ")
}
