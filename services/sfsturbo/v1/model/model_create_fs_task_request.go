package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFsTaskRequest Request Object
type CreateFsTaskRequest struct {

	// MIME类型, application/json
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	// 任务类型。例，DU任务取值为dir-usage
	Feature string `json:"feature"`

	Body *FsDirReq `json:"body,omitempty"`
}

func (o CreateFsTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFsTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateFsTaskRequest", string(data)}, " ")
}
