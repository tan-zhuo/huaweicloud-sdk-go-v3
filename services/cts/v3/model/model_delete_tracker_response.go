package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTrackerResponse Response Object
type DeleteTrackerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTrackerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrackerResponse struct{}"
	}

	return strings.Join([]string{"DeleteTrackerResponse", string(data)}, " ")
}
