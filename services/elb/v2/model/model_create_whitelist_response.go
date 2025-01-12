package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWhitelistResponse Response Object
type CreateWhitelistResponse struct {
	Whitelist      *WhitelistResp `json:"whitelist,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o CreateWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWhitelistResponse struct{}"
	}

	return strings.Join([]string{"CreateWhitelistResponse", string(data)}, " ")
}
