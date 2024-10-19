package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTrafficMirrorFilterOption
type UpdateTrafficMirrorFilterOption struct {

	// 功能说明：流量镜像筛选条件的描述信息 取值范围：0-255个字符，不能包含“<”和“>”
	Description *string `json:"description,omitempty"`

	// 功能说明：流量镜像筛选条件的名称 取值范围：1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`
}

func (o UpdateTrafficMirrorFilterOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTrafficMirrorFilterOption struct{}"
	}

	return strings.Join([]string{"UpdateTrafficMirrorFilterOption", string(data)}, " ")
}
