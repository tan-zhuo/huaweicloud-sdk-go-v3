package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGaussMySqlBackupPolicyResponse Response Object
type ShowGaussMySqlBackupPolicyResponse struct {
	BackupPolicy   *BackupPolicy `json:"backup_policy,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowGaussMySqlBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlBackupPolicyResponse", string(data)}, " ")
}
