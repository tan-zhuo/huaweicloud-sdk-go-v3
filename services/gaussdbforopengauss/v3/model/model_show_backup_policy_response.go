package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupPolicyResponse Response Object
type ShowBackupPolicyResponse struct {
	BackupPolicy   *ShowBackupPolicy `json:"backup_policy,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupPolicyResponse", string(data)}, " ")
}
