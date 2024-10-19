package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStackInstanceDeprecatedResponse Response Object
type DeleteStackInstanceDeprecatedResponse struct {

	// 资源栈集操作（stack_set_operation）的唯一Id。  此ID由资源编排服务在生成资源栈集操作的时候生成，为UUID。
	StackSetOperationId *string `json:"stack_set_operation_id,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o DeleteStackInstanceDeprecatedResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStackInstanceDeprecatedResponse struct{}"
	}

	return strings.Join([]string{"DeleteStackInstanceDeprecatedResponse", string(data)}, " ")
}
