package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMbTasksReportResponse Response Object
type CreateMbTasksReportResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateMbTasksReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMbTasksReportResponse struct{}"
	}

	return strings.Join([]string{"CreateMbTasksReportResponse", string(data)}, " ")
}
