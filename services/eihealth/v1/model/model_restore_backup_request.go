package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreBackupRequest Request Object
type RestoreBackupRequest struct {

	// 归档ID
	BackupId string `json:"backup_id"`

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *RestoreBackupReq `json:"body,omitempty"`
}

func (o RestoreBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreBackupRequest struct{}"
	}

	return strings.Join([]string{"RestoreBackupRequest", string(data)}, " ")
}
