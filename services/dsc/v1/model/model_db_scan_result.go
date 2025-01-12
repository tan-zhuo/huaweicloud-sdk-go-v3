package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DbScanResult 数据库扫描结果
type DbScanResult struct {

	// 扫描结果总数
	Total *int32 `json:"total,omitempty"`

	// 数据库扫描结果列表
	DbScanResults *[]DbScanResultInfo `json:"db_scan_results,omitempty"`
}

func (o DbScanResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbScanResult struct{}"
	}

	return strings.Join([]string{"DbScanResult", string(data)}, " ")
}
