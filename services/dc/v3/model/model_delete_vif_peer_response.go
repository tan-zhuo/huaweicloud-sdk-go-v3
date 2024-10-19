package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVifPeerResponse Response Object
type DeleteVifPeerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVifPeerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVifPeerResponse struct{}"
	}

	return strings.Join([]string{"DeleteVifPeerResponse", string(data)}, " ")
}
