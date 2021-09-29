package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ChangeGaussMySqlInstanceSpecificationRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，严格匹配UUID规则。

	InstanceId string `json:"instance_id"`

	Body *MysqlChangeSpecificationRequest `json:"body,omitempty"`
}

func (o ChangeGaussMySqlInstanceSpecificationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeGaussMySqlInstanceSpecificationRequest struct{}"
	}

	return strings.Join([]string{"ChangeGaussMySqlInstanceSpecificationRequest", string(data)}, " ")
}