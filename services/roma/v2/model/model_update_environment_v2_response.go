package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnvironmentV2Response Response Object
type UpdateEnvironmentV2Response struct {

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 环境名称
	Name *string `json:"name,omitempty"`

	// 描述信息
	Remark *string `json:"remark,omitempty"`

	// 环境编号
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEnvironmentV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnvironmentV2Response struct{}"
	}

	return strings.Join([]string{"UpdateEnvironmentV2Response", string(data)}, " ")
}
