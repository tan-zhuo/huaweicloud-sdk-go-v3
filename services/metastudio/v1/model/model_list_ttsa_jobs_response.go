package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTtsaJobsResponse Response Object
type ListTtsaJobsResponse struct {

	// 语音驱动任务总数。
	Total *int32 `json:"total,omitempty"`

	// 语音驱动任务列表。
	TtsaJobs *[]TtsaJob `json:"ttsa_jobs,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListTtsaJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTtsaJobsResponse struct{}"
	}

	return strings.Join([]string{"ListTtsaJobsResponse", string(data)}, " ")
}
