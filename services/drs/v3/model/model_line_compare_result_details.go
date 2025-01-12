package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LineCompareResultDetails struct {

	// 源库名称。
	SourceDbName string `json:"source_db_name"`

	// 该库的表的行对比详情。
	LineCompareDetail []LineCompareDetail `json:"line_compare_detail"`

	// 该库的行对比结果详情总数。
	LineCompareDetailCount int32 `json:"line_compare_detail_count"`
}

func (o LineCompareResultDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineCompareResultDetails struct{}"
	}

	return strings.Join([]string{"LineCompareResultDetails", string(data)}, " ")
}
