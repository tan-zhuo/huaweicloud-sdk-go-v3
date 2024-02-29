package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactoryPackageDetailRequest Request Object
type ShowFactoryPackageDetailRequest struct {

	// 发布包ID
	PackageId string `json:"package_id"`

	// 工作空间ID，默认查询default
	Workspace *string `json:"workspace,omitempty"`

	// 有Body体的情况下必须，无Body体的情况下则无需填写和校验，默认值：application/json
	ContentType *string `json:"Content-Type,omitempty"`
}

func (o ShowFactoryPackageDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactoryPackageDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowFactoryPackageDetailRequest", string(data)}, " ")
}
