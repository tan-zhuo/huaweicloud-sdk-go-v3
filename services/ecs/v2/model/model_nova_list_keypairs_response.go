package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaListKeypairsResponse Response Object
type NovaListKeypairsResponse struct {

	// 密钥信息列表。
	Keypairs       *[]NovaListKeypairsResult `json:"keypairs,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o NovaListKeypairsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListKeypairsResponse struct{}"
	}

	return strings.Join([]string{"NovaListKeypairsResponse", string(data)}, " ")
}
