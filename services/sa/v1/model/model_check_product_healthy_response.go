package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckProductHealthyResponse Response Object
type CheckProductHealthyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CheckProductHealthyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckProductHealthyResponse struct{}"
	}

	return strings.Join([]string{"CheckProductHealthyResponse", string(data)}, " ")
}
