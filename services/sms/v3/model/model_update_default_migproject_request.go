package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDefaultMigprojectRequest Request Object
type UpdateDefaultMigprojectRequest struct {

	// 迁移项目ID
	MigProjectId string `json:"mig_project_id"`
}

func (o UpdateDefaultMigprojectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDefaultMigprojectRequest struct{}"
	}

	return strings.Join([]string{"UpdateDefaultMigprojectRequest", string(data)}, " ")
}
