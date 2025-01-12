package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNaRequest Request Object
type ShowNaRequest struct {

	// 北向数据接收端点ID
	NaId string `json:"na_id"`
}

func (o ShowNaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNaRequest struct{}"
	}

	return strings.Join([]string{"ShowNaRequest", string(data)}, " ")
}
