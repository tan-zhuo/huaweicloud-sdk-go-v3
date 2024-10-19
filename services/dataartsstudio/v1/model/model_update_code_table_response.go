package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCodeTableResponse Response Object
type UpdateCodeTableResponse struct {
	Data           *CreateCodeTableResultData `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o UpdateCodeTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCodeTableResponse struct{}"
	}

	return strings.Join([]string{"UpdateCodeTableResponse", string(data)}, " ")
}
