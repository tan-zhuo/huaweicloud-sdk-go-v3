package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FinancialStatementResultImageSize 图像信息
type FinancialStatementResultImageSize struct {

	// 矫正后图像的高。
	Height *int32 `json:"height,omitempty"`

	// 矫正后图像的宽。
	Width *int32 `json:"width,omitempty"`
}

func (o FinancialStatementResultImageSize) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FinancialStatementResultImageSize struct{}"
	}

	return strings.Join([]string{"FinancialStatementResultImageSize", string(data)}, " ")
}
