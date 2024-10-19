package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackSetOperationCreateTimePrimitiveTypeHolder struct {

	// 资源栈集操作的创建时间，格式为YYYY-MM-DDTHH:mm:ss.SSSZ，精确到毫秒，UTC时区，即，如1970-01-01T00:00:00.000Z。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o StackSetOperationCreateTimePrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSetOperationCreateTimePrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackSetOperationCreateTimePrimitiveTypeHolder", string(data)}, " ")
}
