package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScriptJobsResponse Response Object
type ListScriptJobsResponse struct {
	Data           *JobScriptOrderListPage `json:"data,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListScriptJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptJobsResponse struct{}"
	}

	return strings.Join([]string{"ListScriptJobsResponse", string(data)}, " ")
}
