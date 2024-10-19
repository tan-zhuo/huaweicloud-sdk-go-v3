package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainParseDetailResponse Response Object
type ListDomainParseDetailResponse struct {

	// 域名id列表
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDomainParseDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainParseDetailResponse struct{}"
	}

	return strings.Join([]string{"ListDomainParseDetailResponse", string(data)}, " ")
}
