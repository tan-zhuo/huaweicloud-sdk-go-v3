package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWhitelistsResponse Response Object
type ListWhitelistsResponse struct {

	// 白名单对象的列表
	Whitelists     *[]WhitelistResp `json:"whitelists,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListWhitelistsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWhitelistsResponse struct{}"
	}

	return strings.Join([]string{"ListWhitelistsResponse", string(data)}, " ")
}
