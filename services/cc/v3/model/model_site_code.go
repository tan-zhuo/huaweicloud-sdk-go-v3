package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteCode 站点编码
type SiteCode struct {

	// 站点编码定义
	SiteCode string `json:"site_code"`
}

func (o SiteCode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteCode struct{}"
	}

	return strings.Join([]string{"SiteCode", string(data)}, " ")
}
