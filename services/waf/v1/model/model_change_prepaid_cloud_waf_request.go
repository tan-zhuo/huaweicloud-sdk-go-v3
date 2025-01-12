package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangePrepaidCloudWafRequest Request Object
type ChangePrepaidCloudWafRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *ChangePrepaidCloudWafRequestBody `json:"body,omitempty"`
}

func (o ChangePrepaidCloudWafRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangePrepaidCloudWafRequest struct{}"
	}

	return strings.Join([]string{"ChangePrepaidCloudWafRequest", string(data)}, " ")
}
