package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppQuotaRequest Request Object
type UpdateAppQuotaRequest struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 凭据配额编号
	AppQuotaId string `json:"app_quota_id"`

	Body *AppQuotaCreate `json:"body,omitempty"`
}

func (o UpdateAppQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppQuotaRequest struct{}"
	}

	return strings.Join([]string{"UpdateAppQuotaRequest", string(data)}, " ")
}
