package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiBatchPublish struct {

	// 需要发布或下线的API ID列表，单次更新上限为1000个API。必须指定apis或group_id。
	Apis *[]string `json:"apis,omitempty"`

	// 环境ID
	EnvId string `json:"env_id"`

	// API分组ID。必须指定apis或group_id。
	GroupId *string `json:"group_id,omitempty"`

	// 对本次发布的描述信息  字符长度不超过255 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`
}

func (o ApiBatchPublish) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiBatchPublish struct{}"
	}

	return strings.Join([]string{"ApiBatchPublish", string(data)}, " ")
}
