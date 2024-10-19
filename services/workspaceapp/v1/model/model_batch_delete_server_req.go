package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteServerReq 批量删除服务器请求。
type BatchDeleteServerReq struct {

	// 批量请求的服务器ID列表，一次请求数量区间 [1, 20]。
	Items []string `json:"items"`
}

func (o BatchDeleteServerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteServerReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteServerReq", string(data)}, " ")
}
