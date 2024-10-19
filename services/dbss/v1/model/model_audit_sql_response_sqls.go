package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditSqlResponseSqls struct {
	Sql *AuditSqlResponseSql `json:"sql"`
}

func (o AuditSqlResponseSqls) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditSqlResponseSqls struct{}"
	}

	return strings.Join([]string{"AuditSqlResponseSqls", string(data)}, " ")
}
