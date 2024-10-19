package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListScriptJobBatchesRequest Request Object
type ListScriptJobBatchesRequest struct {

	// 脚本工单的执行Id，取自executeJobScript和ListJobScriptOrders返回体中
	ExecuteUuid string `json:"execute_uuid"`

	// 国际化标记，zh-cn表示中文，en-us或不传表示英文
	XLanguage *string `json:"X-Language,omitempty"`

	// 项目ID，一个项目对应一个region
	XProjectId *string `json:"x-project-id,omitempty"`

	// IAM5.0用户信息
	XUserProfile *string `json:"x-user-profile,omitempty"`
}

func (o ListScriptJobBatchesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptJobBatchesRequest struct{}"
	}

	return strings.Join([]string{"ListScriptJobBatchesRequest", string(data)}, " ")
}
