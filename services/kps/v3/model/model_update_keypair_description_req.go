package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeypairDescriptionReq 更新SSH密钥对描述消息体
type UpdateKeypairDescriptionReq struct {

	// 描述信息
	Description string `json:"description"`
}

func (o UpdateKeypairDescriptionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeypairDescriptionReq struct{}"
	}

	return strings.Join([]string{"UpdateKeypairDescriptionReq", string(data)}, " ")
}
