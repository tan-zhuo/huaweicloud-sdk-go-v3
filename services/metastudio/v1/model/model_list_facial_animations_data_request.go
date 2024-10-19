package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFacialAnimationsDataRequest Request Object
type ListFacialAnimationsDataRequest struct {

	// 表情驱动任务ID
	JobId string `json:"job_id"`
}

func (o ListFacialAnimationsDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFacialAnimationsDataRequest struct{}"
	}

	return strings.Join([]string{"ListFacialAnimationsDataRequest", string(data)}, " ")
}
