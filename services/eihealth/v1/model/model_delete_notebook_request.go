package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNotebookRequest Request Object
type DeleteNotebookRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// notebook id
	NotebookId string `json:"notebook_id"`
}

func (o DeleteNotebookRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotebookRequest struct{}"
	}

	return strings.Join([]string{"DeleteNotebookRequest", string(data)}, " ")
}
