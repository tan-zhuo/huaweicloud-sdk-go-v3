package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccessConfigRequestBody 修改日志接入请求体
type UpdateAccessConfigRequestBody struct {

	// 日志接入ID
	AccessConfigId string `json:"access_config_id"`

	AccessConfigDetail *AccessConfigDeatilUpdate `json:"access_config_detail,omitempty"`

	HostGroupInfo *AccessConfigHostGroupIdList `json:"host_group_info,omitempty"`

	// 标签信息。KEY不能重复,最多20个标签
	AccessConfigTag *[]AccessConfigTag `json:"access_config_tag,omitempty"`

	// 日志拆分
	LogSplit *bool `json:"log_split,omitempty"`

	// 二进制采集
	BinaryCollect *bool `json:"binary_collect,omitempty"`

	// CCE集群ID，CCE类型时，为必填
	ClusterId *string `json:"cluster_id,omitempty"`
}

func (o UpdateAccessConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccessConfigRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAccessConfigRequestBody", string(data)}, " ")
}
