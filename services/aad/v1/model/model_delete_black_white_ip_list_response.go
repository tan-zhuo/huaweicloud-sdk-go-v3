package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBlackWhiteIpListResponse Response Object
type DeleteBlackWhiteIpListResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBlackWhiteIpListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBlackWhiteIpListResponse struct{}"
	}

	return strings.Join([]string{"DeleteBlackWhiteIpListResponse", string(data)}, " ")
}
