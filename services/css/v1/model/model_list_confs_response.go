package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfsResponse Response Object
type ListConfsResponse struct {

	// 配置文件列表。
	Confs          *[]Confs `json:"confs,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListConfsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfsResponse struct{}"
	}

	return strings.Join([]string{"ListConfsResponse", string(data)}, " ")
}
