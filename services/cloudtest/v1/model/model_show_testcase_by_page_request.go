package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestcaseByPageRequest Request Object
type ShowTestcaseByPageRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CommRequestTestCasePageParam `json:"body,omitempty"`
}

func (o ShowTestcaseByPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestcaseByPageRequest struct{}"
	}

	return strings.Join([]string{"ShowTestcaseByPageRequest", string(data)}, " ")
}
