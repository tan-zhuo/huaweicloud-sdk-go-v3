package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogtanksResponse Response Object
type ListLogtanksResponse struct {

	// 描述信息
	Logtanks *[]Logtank `json:"logtanks,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLogtanksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogtanksResponse struct{}"
	}

	return strings.Join([]string{"ListLogtanksResponse", string(data)}, " ")
}
