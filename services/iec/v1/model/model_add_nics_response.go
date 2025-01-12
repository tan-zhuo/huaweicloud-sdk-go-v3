package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddNicsResponse Response Object
type AddNicsResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddNicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddNicsResponse struct{}"
	}

	return strings.Join([]string{"AddNicsResponse", string(data)}, " ")
}
