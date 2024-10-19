package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AllowGuestUnmuteResponse Response Object
type AllowGuestUnmuteResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AllowGuestUnmuteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllowGuestUnmuteResponse struct{}"
	}

	return strings.Join([]string{"AllowGuestUnmuteResponse", string(data)}, " ")
}
