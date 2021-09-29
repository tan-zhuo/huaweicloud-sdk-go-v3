package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateGaussMySqlInstanceNameRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，严格匹配UUID规则。

	InstanceId string `json:"instance_id"`

	Body *MysqlUpdateInstanceNameRequest `json:"body,omitempty"`
}

func (o UpdateGaussMySqlInstanceNameRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlInstanceNameRequest", string(data)}, " ")
}