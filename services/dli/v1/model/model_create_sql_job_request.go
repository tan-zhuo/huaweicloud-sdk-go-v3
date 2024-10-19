package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSqlJobRequest Request Object
type CreateSqlJobRequest struct {
	Body *CreateSqlJobRequestBody `json:"body,omitempty"`
}

func (o CreateSqlJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlJobRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlJobRequest", string(data)}, " ")
}
