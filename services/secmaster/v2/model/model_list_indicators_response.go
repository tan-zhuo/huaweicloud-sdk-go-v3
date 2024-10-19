package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIndicatorsResponse Response Object
type ListIndicatorsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 指标列表数据
	Data *[]IndicatorDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListIndicatorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIndicatorsResponse struct{}"
	}

	return strings.Join([]string{"ListIndicatorsResponse", string(data)}, " ")
}
