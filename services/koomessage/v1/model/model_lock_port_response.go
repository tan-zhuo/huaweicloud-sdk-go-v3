package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LockPortResponse Response Object
type LockPortResponse struct {
	Data           *LockPortResponseModel `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o LockPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockPortResponse struct{}"
	}

	return strings.Join([]string{"LockPortResponse", string(data)}, " ")
}
