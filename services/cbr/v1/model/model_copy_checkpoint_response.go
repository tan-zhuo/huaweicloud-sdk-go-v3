package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyCheckpointResponse Response Object
type CopyCheckpointResponse struct {
	Replication    *CheckpointReplicateRespBody `json:"replication,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o CopyCheckpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyCheckpointResponse struct{}"
	}

	return strings.Join([]string{"CopyCheckpointResponse", string(data)}, " ")
}
