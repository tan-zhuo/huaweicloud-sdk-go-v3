package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLinkResponse Response Object
type DeleteLinkResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLinkResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLinkResponse struct{}"
	}

	return strings.Join([]string{"DeleteLinkResponse", string(data)}, " ")
}
