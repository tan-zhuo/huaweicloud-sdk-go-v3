package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWhitelistResponse Response Object
type ShowWhitelistResponse struct {
	Whitelist      *WhitelistResp `json:"whitelist,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWhitelistResponse struct{}"
	}

	return strings.Join([]string{"ShowWhitelistResponse", string(data)}, " ")
}
