package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApiGroupsV2Response Response Object
type ListApiGroupsV2Response struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 分组列表
	Groups         *[]ApiGroupInfo `json:"groups,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListApiGroupsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiGroupsV2Response struct{}"
	}

	return strings.Join([]string{"ListApiGroupsV2Response", string(data)}, " ")
}
