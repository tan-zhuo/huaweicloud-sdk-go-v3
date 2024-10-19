package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyHarvestTaskRequest Request Object
type ModifyHarvestTaskRequest struct {

	// 服务鉴权Token，服务开启鉴权，必须携带Access-Control-Allow-Internal访问服务。
	AccessControlAllowInternal *string `json:"Access-Control-Allow-Internal,omitempty"`

	// 服务鉴权Token，服务开启鉴权，必须携带Access-Control-Allow-External访问服务。
	AccessControlAllowExternal *string `json:"Access-Control-Allow-External,omitempty"`

	Body *ModifyHarvestTaskRequestBody `json:"body,omitempty"`
}

func (o ModifyHarvestTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyHarvestTaskRequest struct{}"
	}

	return strings.Join([]string{"ModifyHarvestTaskRequest", string(data)}, " ")
}
