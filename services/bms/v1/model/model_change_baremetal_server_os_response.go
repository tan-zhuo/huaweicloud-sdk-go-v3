package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeBaremetalServerOsResponse Response Object
type ChangeBaremetalServerOsResponse struct {

	// 提交任务成功后返回的任务ID，用户可以使用该ID对任务执行情况进行查询
	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeBaremetalServerOsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeBaremetalServerOsResponse struct{}"
	}

	return strings.Join([]string{"ChangeBaremetalServerOsResponse", string(data)}, " ")
}
