package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEnvironmentV2Response Response Object
type CreateEnvironmentV2Response struct {

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 环境名称
	Name *string `json:"name,omitempty"`

	// 描述信息
	Remark *string `json:"remark,omitempty"`

	// 环境id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEnvironmentV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentV2Response struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentV2Response", string(data)}, " ")
}
