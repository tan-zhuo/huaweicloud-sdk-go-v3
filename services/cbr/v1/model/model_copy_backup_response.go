package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyBackupResponse Response Object
type CopyBackupResponse struct {
	Replication    *BackupReplicateRespBody `json:"replication,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CopyBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyBackupResponse struct{}"
	}

	return strings.Join([]string{"CopyBackupResponse", string(data)}, " ")
}
