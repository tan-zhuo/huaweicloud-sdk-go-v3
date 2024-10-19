package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBlackWhiteListResponse Response Object
type UpdateBlackWhiteListResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateBlackWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBlackWhiteListResponse struct{}"
	}

	return strings.Join([]string{"UpdateBlackWhiteListResponse", string(data)}, " ")
}
