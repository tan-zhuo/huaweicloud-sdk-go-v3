package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotebookRequest Request Object
type UpdateNotebookRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// notebook id
	NotebookId string `json:"notebook_id"`

	Body *UpdateNotebookReq `json:"body,omitempty"`
}

func (o UpdateNotebookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotebookRequest struct{}"
	}

	return strings.Join([]string{"UpdateNotebookRequest", string(data)}, " ")
}
