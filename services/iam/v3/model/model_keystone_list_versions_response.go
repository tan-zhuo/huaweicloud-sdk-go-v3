package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListVersionsResponse Response Object
type KeystoneListVersionsResponse struct {
	Versions       *Versions `json:"versions,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o KeystoneListVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListVersionsResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListVersionsResponse", string(data)}, " ")
}
