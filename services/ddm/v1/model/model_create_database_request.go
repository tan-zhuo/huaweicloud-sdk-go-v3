package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDatabaseRequest Request Object
type CreateDatabaseRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateDatabaseReq `json:"body,omitempty"`
}

func (o CreateDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreateDatabaseRequest", string(data)}, " ")
}
