package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UniqueConstraintInput struct {

	// 列名
	ColumnName string `json:"column_name"`

	// constraint Name
	ConstraintName string `json:"constraint_name"`

	// 列序号（限制条件中第几位）
	KeySequence *int32 `json:"key_sequence,omitempty"`

	// enable constraint
	EnableConstraint bool `json:"enable_constraint"`

	// constraint is rely when Query
	RelyConstraint bool `json:"rely_constraint"`

	// constraint is validated
	ValidateConstraint bool `json:"validate_constraint"`
}

func (o UniqueConstraintInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UniqueConstraintInput struct{}"
	}

	return strings.Join([]string{"UniqueConstraintInput", string(data)}, " ")
}
