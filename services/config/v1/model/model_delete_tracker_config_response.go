package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrackerConfigResponse Response Object
type DeleteTrackerConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTrackerConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrackerConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteTrackerConfigResponse", string(data)}, " ")
}
