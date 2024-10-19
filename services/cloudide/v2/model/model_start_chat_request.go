package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartChatRequest Request Object
type StartChatRequest struct {
	Body *StartChatRequestMessage `json:"body,omitempty"`
}

func (o StartChatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartChatRequest struct{}"
	}

	return strings.Join([]string{"StartChatRequest", string(data)}, " ")
}
