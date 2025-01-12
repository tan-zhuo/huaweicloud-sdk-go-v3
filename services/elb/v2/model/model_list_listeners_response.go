package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListListenersResponse Response Object
type ListListenersResponse struct {

	// 监听器对象列表
	Listeners      *[]ListenerResp `json:"listeners,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListListenersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersResponse struct{}"
	}

	return strings.Join([]string{"ListListenersResponse", string(data)}, " ")
}
