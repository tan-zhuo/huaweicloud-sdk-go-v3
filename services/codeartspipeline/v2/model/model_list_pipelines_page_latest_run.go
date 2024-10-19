package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelinesPageLatestRun 最近一次运行信息
type ListPipelinesPageLatestRun struct {

	// 流水线ID
	PipelineId *string `json:"pipeline_id,omitempty"`

	// 流水线运行实例ID
	PipelineRunId *string `json:"pipeline_run_id,omitempty"`

	// 执行人ID
	ExecutorId *string `json:"executor_id,omitempty"`

	// 执行人名称
	ExecutorName *string `json:"executor_name,omitempty"`

	// 阶段信息列表
	StageStatusList *[]ListPipelinesPageLatestRunStageStatusList `json:"stage_status_list,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 运行序号
	RunNumber *int32 `json:"run_number,omitempty"`

	// 触发类型
	TriggerType *string `json:"trigger_type,omitempty"`

	BuildParams *ListPipelinesPageLatestRunBuildParams `json:"build_params,omitempty"`

	ArtifactParams *ListPipelinesPageLatestRunArtifactParams `json:"artifact_params,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 修改页地址
	ModifyUrl *string `json:"modify_url,omitempty"`

	// 详情页地址
	DetailUrl *string `json:"detail_url,omitempty"`
}

func (o ListPipelinesPageLatestRun) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelinesPageLatestRun struct{}"
	}

	return strings.Join([]string{"ListPipelinesPageLatestRun", string(data)}, " ")
}
