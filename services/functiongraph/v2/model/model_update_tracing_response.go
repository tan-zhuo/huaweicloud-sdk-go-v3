package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTracingResponse Response Object
type UpdateTracingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateTracingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTracingResponse struct{}"
	}

	return strings.Join([]string{"UpdateTracingResponse", string(data)}, " ")
}
