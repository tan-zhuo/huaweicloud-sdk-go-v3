package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddBlackWhiteListResponse Response Object
type AddBlackWhiteListResponse struct {
	Data           *IdObject `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AddBlackWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddBlackWhiteListResponse struct{}"
	}

	return strings.Join([]string{"AddBlackWhiteListResponse", string(data)}, " ")
}
