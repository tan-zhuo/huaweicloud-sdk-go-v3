package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApprovePlaybookInfo 审核详情
type ApprovePlaybookInfo struct {

	// 审核结果  通过：PASS 不通过：UN_PASS
	Result *string `json:"result,omitempty"`

	// 审核意见
	Content *string `json:"content,omitempty"`
}

func (o ApprovePlaybookInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApprovePlaybookInfo struct{}"
	}

	return strings.Join([]string{"ApprovePlaybookInfo", string(data)}, " ")
}
