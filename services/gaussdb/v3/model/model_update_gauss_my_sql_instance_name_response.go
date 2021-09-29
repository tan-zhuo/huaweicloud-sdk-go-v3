package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateGaussMySqlInstanceNameResponse struct {
	// 修改实例名称的任务id

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateGaussMySqlInstanceNameResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstanceNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstanceNameResponse", string(data)}, " ")
}