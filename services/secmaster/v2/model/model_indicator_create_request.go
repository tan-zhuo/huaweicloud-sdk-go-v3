package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndicatorCreateRequest 创建指标请求参数
type IndicatorCreateRequest struct {
	DataObject *CreateIndicatorDetail `json:"data_object"`
}

func (o IndicatorCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorCreateRequest struct{}"
	}

	return strings.Join([]string{"IndicatorCreateRequest", string(data)}, " ")
}
