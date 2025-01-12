package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAclStrategyV2Response Response Object
type UpdateAclStrategyV2Response struct {

	// 名称
	AclName *string `json:"acl_name,omitempty"`

	// 类型: - PERMIT（白名单类型） - DENY（黑名单类型）
	AclType *string `json:"acl_type,omitempty"`

	// ACL策略值
	AclValue *string `json:"acl_value,omitempty"`

	// 对象类型： - IP - DOMAIN [Site自治版不支持DOMAIN类型。](tag:Site)
	EntityType *string `json:"entity_type,omitempty"`

	// 编号
	Id *string `json:"id,omitempty"`

	// 更新时间
	UpdateTime     *sdktime.SdkTime `json:"update_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateAclStrategyV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAclStrategyV2Response struct{}"
	}

	return strings.Join([]string{"UpdateAclStrategyV2Response", string(data)}, " ")
}
