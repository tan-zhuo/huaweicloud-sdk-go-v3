package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommitCheckpointResponse Response Object
type CommitCheckpointResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CommitCheckpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitCheckpointResponse struct{}"
	}

	return strings.Join([]string{"CommitCheckpointResponse", string(data)}, " ")
}
