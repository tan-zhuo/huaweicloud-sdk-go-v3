package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBackupResponse Response Object
type DeleteBackupResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupResponse struct{}"
	}

	return strings.Join([]string{"DeleteBackupResponse", string(data)}, " ")
}
