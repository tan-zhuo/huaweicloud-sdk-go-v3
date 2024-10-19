package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClickHouseDataBaseConfigRequest Request Object
type DeleteClickHouseDataBaseConfigRequest struct {

	// ClickHouse实例ID，严格匹配UUID规则。
	InstanceId string `json:"instance_id"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 数据库名。
	Database string `json:"database"`
}

func (o DeleteClickHouseDataBaseConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClickHouseDataBaseConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteClickHouseDataBaseConfigRequest", string(data)}, " ")
}
