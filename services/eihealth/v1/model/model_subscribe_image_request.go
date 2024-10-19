package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscribeImageRequest Request Object
type SubscribeImageRequest struct {

	// 平台项目ID，您可以在平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *SubscribeImageReq `json:"body,omitempty"`
}

func (o SubscribeImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscribeImageRequest struct{}"
	}

	return strings.Join([]string{"SubscribeImageRequest", string(data)}, " ")
}
