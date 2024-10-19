package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccessControl 访问控制配置，允许配置黑名单、白名单。
type AccessControl struct {

	// 黑名单数组。 - 每行一个IP地址或网段，以回车结束。 - 每个IP地址组最多可添加300个IP地址或网段。
	Black *[]string `json:"black,omitempty"`

	// 白名单数组。 - 每行一个IP地址或网段，以回车结束。 - 每个IP地址组最多可添加300个IP地址或网段。
	White *[]string `json:"white,omitempty"`
}

func (o AccessControl) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessControl struct{}"
	}

	return strings.Join([]string{"AccessControl", string(data)}, " ")
}
