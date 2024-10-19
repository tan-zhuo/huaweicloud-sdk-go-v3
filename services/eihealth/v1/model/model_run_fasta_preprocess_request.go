package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunFastaPreprocessRequest Request Object
type RunFastaPreprocessRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *RunFastaPreprocessReq `json:"body,omitempty"`
}

func (o RunFastaPreprocessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunFastaPreprocessRequest struct{}"
	}

	return strings.Join([]string{"RunFastaPreprocessRequest", string(data)}, " ")
}
