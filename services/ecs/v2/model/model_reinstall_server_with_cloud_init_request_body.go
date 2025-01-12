package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReinstallServerWithCloudInitRequestBody This is a auto create Body Object
type ReinstallServerWithCloudInitRequestBody struct {
	OsReinstall *ReinstallServerWithCloudInitOption `json:"os-reinstall"`
}

func (o ReinstallServerWithCloudInitRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallServerWithCloudInitRequestBody struct{}"
	}

	return strings.Join([]string{"ReinstallServerWithCloudInitRequestBody", string(data)}, " ")
}
