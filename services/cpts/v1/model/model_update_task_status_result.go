package model

import (
	"encoding/json"

	"strings"
)

type UpdateTaskStatusResult struct {
	// task_run_id

	TaskRunId *int32 `json:"task_run_id,omitempty"`
}

func (o UpdateTaskStatusResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTaskStatusResult struct{}"
	}

	return strings.Join([]string{"UpdateTaskStatusResult", string(data)}, " ")
}
