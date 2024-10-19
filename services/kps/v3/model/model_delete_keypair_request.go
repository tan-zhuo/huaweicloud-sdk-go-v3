package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKeypairRequest Request Object
type DeleteKeypairRequest struct {

	// 密钥对名称
	KeypairName string `json:"keypair_name"`
}

func (o DeleteKeypairRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKeypairRequest struct{}"
	}

	return strings.Join([]string{"DeleteKeypairRequest", string(data)}, " ")
}
