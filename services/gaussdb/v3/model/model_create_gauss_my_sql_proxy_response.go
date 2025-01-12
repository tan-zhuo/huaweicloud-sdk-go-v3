package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGaussMySqlProxyResponse Response Object
type CreateGaussMySqlProxyResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGaussMySqlProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlProxyResponse struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlProxyResponse", string(data)}, " ")
}
