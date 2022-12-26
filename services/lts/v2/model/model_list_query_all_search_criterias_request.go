package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListQueryAllSearchCriteriasRequest struct {

	// 租户想查询的日志流所在的日志组的groupid，一般为36位字符串。  缺省值：None  最小长度：36  最大长度：36
	GroupId string `json:"group_id"`
}

func (o ListQueryAllSearchCriteriasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryAllSearchCriteriasRequest struct{}"
	}

	return strings.Join([]string{"ListQueryAllSearchCriteriasRequest", string(data)}, " ")
}
