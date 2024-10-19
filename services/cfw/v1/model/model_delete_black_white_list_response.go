package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBlackWhiteListResponse Response Object
type DeleteBlackWhiteListResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteBlackWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBlackWhiteListResponse struct{}"
	}

	return strings.Join([]string{"DeleteBlackWhiteListResponse", string(data)}, " ")
}
