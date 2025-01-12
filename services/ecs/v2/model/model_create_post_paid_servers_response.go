package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePostPaidServersResponse Response Object
type CreatePostPaidServersResponse struct {

	// 提交任务成功后返回的任务ID，用户可以使用该ID对任务执行情况进行查询。
	JobId *string `json:"job_id,omitempty"`

	// 云服务器ID列表。
	ServerIds      *[]string `json:"serverIds,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreatePostPaidServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidServersResponse struct{}"
	}

	return strings.Join([]string{"CreatePostPaidServersResponse", string(data)}, " ")
}
