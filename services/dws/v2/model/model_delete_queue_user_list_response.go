package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteQueueUserListResponse Response Object
type DeleteQueueUserListResponse struct {

	// 响应编码。
	WorkloadResCode *int32 `json:"workload_res_code,omitempty"`

	// 响应信息。
	WorkloadResStr *string `json:"workload_res_str,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteQueueUserListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueueUserListResponse struct{}"
	}

	return strings.Join([]string{"DeleteQueueUserListResponse", string(data)}, " ")
}
