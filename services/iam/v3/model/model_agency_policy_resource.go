package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyPolicyResource
type AgencyPolicyResource struct {

	// 委托资源的URI。格式为：/iam/agencies/委托ID。例： ``` \"uri\": [\"/iam/agencies/07805acaba800fdd4fbdc00b8f888c7c\"] ```
	Uri []string `json:"uri"`
}

func (o AgencyPolicyResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyPolicyResource struct{}"
	}

	return strings.Join([]string{"AgencyPolicyResource", string(data)}, " ")
}
