package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateReplicationNameResponse Response Object
type UpdateReplicationNameResponse struct {
	Replication    *ShowReplicationParams `json:"replication,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o UpdateReplicationNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReplicationNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateReplicationNameResponse", string(data)}, " ")
}
