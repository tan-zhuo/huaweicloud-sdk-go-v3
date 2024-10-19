package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStackInstanceRequest Request Object
type DeleteStackInstanceRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 资源栈集的名称。此名字在domain_id+region下应唯一，可以使用中文、大小写英文、数字、下划线、中划线。首字符需为中文或者英文，区分大小写。
	StackSetName string `json:"stack_set_name"`

	Body *DeleteStackInstanceRequestBody `json:"body,omitempty"`
}

func (o DeleteStackInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStackInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteStackInstanceRequest", string(data)}, " ")
}
