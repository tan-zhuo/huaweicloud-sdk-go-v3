package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagQuota 配额
type TagQuota struct {

	// 配额键
	QuotaKey string `json:"quota_key"`

	// 配额值
	QuotaLimit int32 `json:"quota_limit"`

	// 已使用
	Used int32 `json:"used"`

	// 单位
	Unit string `json:"unit"`
}

func (o TagQuota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagQuota struct{}"
	}

	return strings.Join([]string{"TagQuota", string(data)}, " ")
}
