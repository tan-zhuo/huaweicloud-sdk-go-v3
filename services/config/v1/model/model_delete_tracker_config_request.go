package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrackerConfigRequest Request Object
type DeleteTrackerConfigRequest struct {
}

func (o DeleteTrackerConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrackerConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteTrackerConfigRequest", string(data)}, " ")
}
