package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchStartInstanceResponse Response Object
type BatchStartInstanceResponse struct {

	// 任务列表对象。
	Jobs           *[]JobResult `json:"jobs,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o BatchStartInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartInstanceResponse struct{}"
	}

	return strings.Join([]string{"BatchStartInstanceResponse", string(data)}, " ")
}
