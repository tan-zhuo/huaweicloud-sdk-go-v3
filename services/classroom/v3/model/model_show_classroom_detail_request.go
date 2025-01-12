package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClassroomDetailRequest Request Object
type ShowClassroomDetailRequest struct {

	// 课堂ID
	ClassroomId string `json:"classroom_id"`
}

func (o ShowClassroomDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClassroomDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowClassroomDetailRequest", string(data)}, " ")
}
