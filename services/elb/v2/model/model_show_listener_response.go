package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowListenerResponse Response Object
type ShowListenerResponse struct {
	Listener       *ListenerResp `json:"listener,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowListenerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListenerResponse struct{}"
	}

	return strings.Join([]string{"ShowListenerResponse", string(data)}, " ")
}
