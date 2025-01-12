package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAssetModelRequest Request Object
type DeleteAssetModelRequest struct {

	// 模型ID
	ModelId string `json:"model_id"`
}

func (o DeleteAssetModelRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAssetModelRequest struct{}"
	}

	return strings.Join([]string{"DeleteAssetModelRequest", string(data)}, " ")
}
