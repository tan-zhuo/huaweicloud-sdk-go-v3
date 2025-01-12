package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteQueueResponse Response Object
type DeleteQueueResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueueResponse struct{}"
	}

	return strings.Join([]string{"DeleteQueueResponse", string(data)}, " ")
}
