package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBackupRequest Request Object
type CreateBackupRequest struct {

	// 图ID。
	GraphId string `json:"graph_id"`
}

func (o CreateBackupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackupRequest struct{}"
	}

	return strings.Join([]string{"CreateBackupRequest", string(data)}, " ")
}
