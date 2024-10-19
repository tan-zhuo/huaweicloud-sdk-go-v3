package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEcnWithVpcResponse Response Object
type UpdateEcnWithVpcResponse struct {

	// 企业连接网络关联虚拟私有云ID
	Id *string `json:"id,omitempty"`

	// 虚拟私有云ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网ID
	SubnetId *string `json:"subnet_id,omitempty"`

	// 本端子网列表
	LocalSubnetList *[]string `json:"local_subnet_list,omitempty"`

	// 对端子网列表
	RemoteSubnetList *[]string `json:"remote_subnet_list,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 创建时间
	CreatedAt      *sdktime.SdkTime `json:"created_at,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateEcnWithVpcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEcnWithVpcResponse struct{}"
	}

	return strings.Join([]string{"UpdateEcnWithVpcResponse", string(data)}, " ")
}
