package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClusterDnsResponse Response Object
type DeleteClusterDnsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteClusterDnsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterDnsResponse struct{}"
	}

	return strings.Join([]string{"DeleteClusterDnsResponse", string(data)}, " ")
}
