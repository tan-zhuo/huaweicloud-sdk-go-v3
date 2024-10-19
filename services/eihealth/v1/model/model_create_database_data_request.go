package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseDataRequest Request Object
type CreateDatabaseDataRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 数据库实例id
	DatabaseId string `json:"database_id"`

	Body *RowDataReq `json:"body,omitempty"`
}

func (o CreateDatabaseDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseDataRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseDataRequest", string(data)}, " ")
}
