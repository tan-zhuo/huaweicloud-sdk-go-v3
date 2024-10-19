package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStackSetResponse Response Object
type UpdateStackSetResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateStackSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStackSetResponse struct{}"
	}

	return strings.Join([]string{"UpdateStackSetResponse", string(data)}, " ")
}
