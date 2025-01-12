package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushTranscriberJobsRequest Request Object
type PushTranscriberJobsRequest struct {

	// 企业项目ID。SIS支持通过企业项目管理（EPS）对不同用户组和用户的资源使用，进行分账。 获取方法：进入“[企业项目管理](https://console.huaweicloud.com/eps/?region=cn-north-4#/projects/list)”页面，单击企业项目名称，在企业项目详情页获取Enterprise-Project-Id（企业项目ID）。 企业项目创建步骤请参见用户指南。 > 说明： 创建企业项目后，在传参时，有以下三类场景。 - 携带正确的ID，正常使用SIS服务，账单归到企业ID对应的企业项目中。 - 携带错误的ID，正常使用SIS服务，账单的企业项目会被分类为“default”。 - 不携带ID，正常使用SIS服务，账单的企业项目会被分类为“default”。
	EnterpriseProjectId *string `json:"Enterprise-Project-Id,omitempty"`

	Body *PostTranscriberJobs `json:"body,omitempty"`
}

func (o PushTranscriberJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushTranscriberJobsRequest struct{}"
	}

	return strings.Join([]string{"PushTranscriberJobsRequest", string(data)}, " ")
}
