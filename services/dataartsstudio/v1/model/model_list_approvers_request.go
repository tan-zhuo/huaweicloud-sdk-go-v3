package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApproversRequest Request Object
type ListApproversRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 审核人姓名。
	ApproverName *string `json:"approver_name,omitempty"`

	// 每页查询条数，即查询Y条数据。默认值50，取值范围[1,100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整，默认值0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListApproversRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApproversRequest struct{}"
	}

	return strings.Join([]string{"ListApproversRequest", string(data)}, " ")
}
