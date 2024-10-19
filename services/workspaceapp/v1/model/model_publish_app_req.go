package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishAppReq 发布应用请求。
type PublishAppReq struct {

	// 发布应用列表(单次最多20个应用)。
	Items []PublishApp `json:"items"`
}

func (o PublishAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishAppReq struct{}"
	}

	return strings.Join([]string{"PublishAppReq", string(data)}, " ")
}
