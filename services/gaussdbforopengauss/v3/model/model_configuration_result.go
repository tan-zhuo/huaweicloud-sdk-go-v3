package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfigurationResult 参数模板信息。
type ConfigurationResult struct {

	// 参数组ID。
	Id string `json:"id"`

	// 参数组名称。
	Name string `json:"name"`

	// 参数组描述。
	Description *string `json:"description,omitempty"`

	// 引擎版本。
	DatastoreVersion string `json:"datastore_version"`

	// 引擎名称。
	DatastoreName string `json:"datastore_name"`

	// 实例类型。
	HaMode string `json:"ha_mode"`

	// 创建时间，格式为\"yyyy-MM-dd HH:mm:ss\"。
	Created string `json:"created"`

	// 更新时间，格式为\"yyyy-MM-dd HH:mm:ss\"。
	Updated string `json:"updated"`

	// 是否是用户自定义参数模板：  - false，表示为系统默认参数模板。 - true，表示为用户自定义参数模板。
	UserDefined bool `json:"user_defined"`
}

func (o ConfigurationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationResult struct{}"
	}

	return strings.Join([]string{"ConfigurationResult", string(data)}, " ")
}
