package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteGaussMySqlInstanceResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGaussMySqlInstanceResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlInstanceResponse", string(data)}, " ")
}
