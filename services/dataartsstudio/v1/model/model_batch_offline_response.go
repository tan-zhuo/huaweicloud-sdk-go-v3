package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchOfflineResponse Response Object
type BatchOfflineResponse struct {
	Data           *BatchOfflineResultData `json:"data,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o BatchOfflineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOfflineResponse struct{}"
	}

	return strings.Join([]string{"BatchOfflineResponse", string(data)}, " ")
}
