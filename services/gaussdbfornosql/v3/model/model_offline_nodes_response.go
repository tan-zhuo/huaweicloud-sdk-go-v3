package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OfflineNodesResponse Response Object
type OfflineNodesResponse struct {

	// 工作流ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o OfflineNodesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OfflineNodesResponse struct{}"
	}

	return strings.Join([]string{"OfflineNodesResponse", string(data)}, " ")
}
