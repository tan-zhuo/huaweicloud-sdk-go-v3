package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsfederationProtocol
type OsfederationProtocol struct {

	// 协议ID。
	Id string `json:"id"`
}

func (o OsfederationProtocol) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsfederationProtocol struct{}"
	}

	return strings.Join([]string{"OsfederationProtocol", string(data)}, " ")
}
