package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IvsStandardByVideoAndNameAndIdResponseBodyResult 调用返回结果。
type IvsStandardByVideoAndNameAndIdResponseBodyResult struct {

	// 子服务名称。
	ServiceName string `json:"service_name"`

	// 成功的结果数量，与resp_data字段对应。
	Count int32 `json:"count"`

	// 请求列表，用于支持批量调用。目前暂时只支持单个数据查询。
	RespData []StandardRespDataByVideoAndNameAndId `json:"resp_data"`
}

func (o IvsStandardByVideoAndNameAndIdResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IvsStandardByVideoAndNameAndIdResponseBodyResult struct{}"
	}

	return strings.Join([]string{"IvsStandardByVideoAndNameAndIdResponseBodyResult", string(data)}, " ")
}
