package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SATA云硬盘类型预留的容量大小，单位为GB，键值对，包含：reserved（预留）、limit（最大）和in_use（已使用）。
type QuotaDetailGigabytesSata struct {
	// 已使用的数量。

	InUse int32 `json:"in_use"`
	// 最大的数量。

	Limit int32 `json:"limit"`
	// 预留属性。

	Reserved int32 `json:"reserved"`
}

func (o QuotaDetailGigabytesSata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaDetailGigabytesSata struct{}"
	}

	return strings.Join([]string{"QuotaDetailGigabytesSata", string(data)}, " ")
}
