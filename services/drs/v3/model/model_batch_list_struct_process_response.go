package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchListStructProcessResponse Response Object
type BatchListStructProcessResponse struct {

	// 批量查询灾备初始化进度返回列表
	Results *[]QueryStructProcessResp `json:"results,omitempty"`

	// 总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchListStructProcessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListStructProcessResponse struct{}"
	}

	return strings.Join([]string{"BatchListStructProcessResponse", string(data)}, " ")
}
