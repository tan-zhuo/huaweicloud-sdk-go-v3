package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateElbListenerResponse Response Object
type CreateElbListenerResponse struct {

	// 负载均衡器id。
	ElbId          *string `json:"elb_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateElbListenerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateElbListenerResponse struct{}"
	}

	return strings.Join([]string{"CreateElbListenerResponse", string(data)}, " ")
}
