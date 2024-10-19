package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteManualBackupRequest Request Object
type DeleteManualBackupRequest struct {

	// 备份ID。
	BackupId string `json:"backup_id"`
}

func (o DeleteManualBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteManualBackupRequest struct{}"
	}

	return strings.Join([]string{"DeleteManualBackupRequest", string(data)}, " ")
}
