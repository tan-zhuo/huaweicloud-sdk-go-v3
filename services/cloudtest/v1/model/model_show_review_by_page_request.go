package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReviewByPageRequest Request Object
type ShowReviewByPageRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CommRequestReviewPageParam `json:"body,omitempty"`
}

func (o ShowReviewByPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReviewByPageRequest struct{}"
	}

	return strings.Join([]string{"ShowReviewByPageRequest", string(data)}, " ")
}
