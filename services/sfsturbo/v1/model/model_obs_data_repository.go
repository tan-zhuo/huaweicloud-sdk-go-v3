package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ObsDataRepository OBS类型后端存储
type ObsDataRepository struct {

	// OBS桶名称
	Bucket string `json:"bucket"`

	// OBS桶所在的区域域名
	Endpoint string `json:"endpoint"`
}

func (o ObsDataRepository) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObsDataRepository struct{}"
	}

	return strings.Join([]string{"ObsDataRepository", string(data)}, " ")
}
