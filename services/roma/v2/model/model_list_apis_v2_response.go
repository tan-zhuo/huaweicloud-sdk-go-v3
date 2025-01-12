package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApisV2Response Response Object
type ListApisV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次查询到的API列表
	Apis           *[]ApiInfoPerPage `json:"apis,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListApisV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisV2Response struct{}"
	}

	return strings.Join([]string{"ListApisV2Response", string(data)}, " ")
}
