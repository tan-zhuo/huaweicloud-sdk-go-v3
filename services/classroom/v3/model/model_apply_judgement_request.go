package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyJudgementRequest Request Object
type ApplyJudgementRequest struct {
	Body *JudgementTaskRequestBody `json:"body,omitempty"`
}

func (o ApplyJudgementRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyJudgementRequest struct{}"
	}

	return strings.Join([]string{"ApplyJudgementRequest", string(data)}, " ")
}
