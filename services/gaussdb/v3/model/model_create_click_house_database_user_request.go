package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClickHouseDatabaseUserRequest Request Object
type CreateClickHouseDatabaseUserRequest struct {

	// ClickHouse实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ClickHouseDatabaseUserInfo `json:"body,omitempty"`
}

func (o CreateClickHouseDatabaseUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClickHouseDatabaseUserRequest struct{}"
	}

	return strings.Join([]string{"CreateClickHouseDatabaseUserRequest", string(data)}, " ")
}
