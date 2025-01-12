package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetHostViewResponse Response Object
type SetHostViewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetHostViewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetHostViewResponse struct{}"
	}

	return strings.Join([]string{"SetHostViewResponse", string(data)}, " ")
}
