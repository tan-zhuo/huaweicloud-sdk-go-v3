package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScheduledTasksRecordsDetailsResponse Response Object
type ListScheduledTasksRecordsDetailsResponse struct {

	// 定时任务执行记录详情列表。
	TasksRecordsDetails *[]ScheduledTasksRecordsDetails `json:"tasks_records_details,omitempty"`

	// 总个数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListScheduledTasksRecordsDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTasksRecordsDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListScheduledTasksRecordsDetailsResponse", string(data)}, " ")
}
