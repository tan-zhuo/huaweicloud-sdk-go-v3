package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobInfo 构建工程参数。
type JobInfo struct {

	// 创建者。
	CreatedBy *string `json:"CREATED_BY,omitempty"`

	// 执行状态。
	ExecutionStatus *JobInfoExecutionStatus `json:"EXECUTION_STATUS,omitempty"`

	// 工作描述。
	JobDesc *string `json:"JOB_DESC,omitempty"`

	// 工作ID。
	JobId *string `json:"JOB_ID,omitempty"`

	// 工作名称。
	JobName *string `json:"JOB_NAME,omitempty"`

	// 类别。
	JobType *string `json:"JOB_TYPE,omitempty"`

	// 排序ID。
	OrderId *string `json:"ORDER_ID,omitempty"`

	// 创建租户的项目ID。
	ProjectId *string `json:"PROJECT_ID,omitempty"`

	// 实例ID。
	ServiceInstanceId *string `json:"SERVICE_INSTANCE_ID,omitempty"`
}

func (o JobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInfo struct{}"
	}

	return strings.Join([]string{"JobInfo", string(data)}, " ")
}

type JobInfoExecutionStatus struct {
	value string
}

type JobInfoExecutionStatusEnum struct {
	RUNNING   JobInfoExecutionStatus
	FAILED    JobInfoExecutionStatus
	SUCCEEDED JobInfoExecutionStatus
}

func GetJobInfoExecutionStatusEnum() JobInfoExecutionStatusEnum {
	return JobInfoExecutionStatusEnum{
		RUNNING: JobInfoExecutionStatus{
			value: "RUNNING",
		},
		FAILED: JobInfoExecutionStatus{
			value: "FAILED",
		},
		SUCCEEDED: JobInfoExecutionStatus{
			value: "SUCCEEDED",
		},
	}
}

func (c JobInfoExecutionStatus) Value() string {
	return c.value
}

func (c JobInfoExecutionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoExecutionStatus) UnmarshalJSON(b []byte) error {
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
