package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUrlTaskInfoRequest Request Object
type ShowUrlTaskInfoRequest struct {

	// 查询起始时间戳（毫秒），不传默认当天00:00，需与结束时间戳同时指定，时间跨度不能超过24小时。
	StartTime *int64 `json:"start_time,omitempty"`

	// 查询结束时间戳（毫秒），不传默认次日00:00，需与开始时间戳同时指定，时间跨度不能超过24小时。
	EndTime *int64 `json:"end_time,omitempty"`

	// 偏移量：特定数据字段与起始数据字段位置的距离，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 单次查询数据条数，上限为100，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 刷新预热url。
	Url *string `json:"url,omitempty"`

	// 任务类型，REFRESH：刷新任务；PREHEATING：预热任务。
	TaskType *string `json:"task_type,omitempty"`

	// url状态，状态类型：processing：处理中；succeed：完成；failed：失败；waiting：等待；refreshing：刷新中; preheating : 预热中。
	Status *string `json:"status,omitempty"`

	// 文件类型，file:文件;directory:目录。
	FileType *string `json:"file_type,omitempty"`
}

func (o ShowUrlTaskInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUrlTaskInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowUrlTaskInfoRequest", string(data)}, " ")
}
