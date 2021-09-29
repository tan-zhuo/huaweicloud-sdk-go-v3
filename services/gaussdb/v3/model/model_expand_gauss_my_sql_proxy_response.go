package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ExpandGaussMySqlProxyResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandGaussMySqlProxyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExpandGaussMySqlProxyResponse struct{}"
	}

	return strings.Join([]string{"ExpandGaussMySqlProxyResponse", string(data)}, " ")
}
