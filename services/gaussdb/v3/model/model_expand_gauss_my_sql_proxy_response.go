package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExpandGaussMySqlProxyResponse Response Object
type ExpandGaussMySqlProxyResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandGaussMySqlProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandGaussMySqlProxyResponse struct{}"
	}

	return strings.Join([]string{"ExpandGaussMySqlProxyResponse", string(data)}, " ")
}
