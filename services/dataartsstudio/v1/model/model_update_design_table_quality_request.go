package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDesignTableQualityRequest Request Object
type UpdateDesignTableQualityRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 表的ID，填写String类型替代Long类型。
	TableId string `json:"table_id"`

	Body *UpdateDesignTableQualityParam `json:"body,omitempty"`
}

func (o UpdateDesignTableQualityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDesignTableQualityRequest struct{}"
	}

	return strings.Join([]string{"UpdateDesignTableQualityRequest", string(data)}, " ")
}
