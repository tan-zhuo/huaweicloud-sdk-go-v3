package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateNodePoolStatus
type CreateNodePoolStatus struct {

	// 当前节点池中所有节点数量（不含删除中的节点）。
	CurrentNode *int32 `json:"currentNode,omitempty"`

	// 当前节点池中处于创建流程中的节点数量。
	CreatingNode *int32 `json:"creatingNode,omitempty"`

	// 当前节点池中删除中的节点数量。
	DeletingNode *int32 `json:"deletingNode,omitempty"`

	// 节点池状态。 - 空值：可用（节点池当前节点数已达到预期，且无伸缩中的节点） - Synchronizing：伸缩中（节点池当前节点数未达到预期，且无伸缩中的节点） - Synchronized：伸缩等待中（节点池当前节点数未达到预期，或者存在伸缩中的节点） - SoldOut：节点池当前不可扩容（兼容字段，标记节点池资源售罄、资源配额不足等不可扩容状态） > 上述节点池状态已废弃，仅兼容保留，不建议使用，替代感知方式如下： > - 节点池扩缩状态：可通过currentNode/creatingNode/deletingNode节点状态统计信息，精确感知当前节点池扩缩状态。 > - 节点池可扩容状态：可通过conditions感知节点池详细状态，其中\"Scalable\"可替代SoldOut语义。 - Deleting：删除中 - Error：错误
	Phase *CreateNodePoolStatusPhase `json:"phase,omitempty"`

	// 节点池当前详细状态列表，详情参见Condition类型定义。
	Conditions *[]NodePoolCondition `json:"conditions,omitempty"`

	// 伸缩组当前详细状态信息，详情参见ScaleGroupStatus类型定义
	ScaleGroupStatuses *[]ScaleGroupStatus `json:"scaleGroupStatuses,omitempty"`
}

func (o CreateNodePoolStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodePoolStatus struct{}"
	}

	return strings.Join([]string{"CreateNodePoolStatus", string(data)}, " ")
}

type CreateNodePoolStatusPhase struct {
	value string
}

type CreateNodePoolStatusPhaseEnum struct {
	SYNCHRONIZING CreateNodePoolStatusPhase
	SYNCHRONIZED  CreateNodePoolStatusPhase
	SOLD_OUT      CreateNodePoolStatusPhase
	DELETING      CreateNodePoolStatusPhase
	ERROR         CreateNodePoolStatusPhase
}

func GetCreateNodePoolStatusPhaseEnum() CreateNodePoolStatusPhaseEnum {
	return CreateNodePoolStatusPhaseEnum{
		SYNCHRONIZING: CreateNodePoolStatusPhase{
			value: "Synchronizing",
		},
		SYNCHRONIZED: CreateNodePoolStatusPhase{
			value: "Synchronized",
		},
		SOLD_OUT: CreateNodePoolStatusPhase{
			value: "SoldOut",
		},
		DELETING: CreateNodePoolStatusPhase{
			value: "Deleting",
		},
		ERROR: CreateNodePoolStatusPhase{
			value: "Error",
		},
	}
}

func (c CreateNodePoolStatusPhase) Value() string {
	return c.value
}

func (c CreateNodePoolStatusPhase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNodePoolStatusPhase) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
