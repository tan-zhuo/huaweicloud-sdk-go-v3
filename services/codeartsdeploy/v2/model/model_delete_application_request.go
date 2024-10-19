package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApplicationRequest Request Object
type DeleteApplicationRequest struct {

	// 应用id
	AppId string `json:"app_id"`
}

func (o DeleteApplicationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationRequest struct{}"
	}

	return strings.Join([]string{"DeleteApplicationRequest", string(data)}, " ")
}
