package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCodeTableResponse Response Object
type CreateCodeTableResponse struct {
	Data           *CreateCodeTableResultData `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o CreateCodeTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCodeTableResponse struct{}"
	}

	return strings.Join([]string{"CreateCodeTableResponse", string(data)}, " ")
}
