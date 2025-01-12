package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSingleJobExeResponse Response Object
type ShowSingleJobExeResponse struct {
	JobDetail      *JobQueryBean `json:"job_detail,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowSingleJobExeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSingleJobExeResponse struct{}"
	}

	return strings.Join([]string{"ShowSingleJobExeResponse", string(data)}, " ")
}
