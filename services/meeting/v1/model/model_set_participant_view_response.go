package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetParticipantViewResponse Response Object
type SetParticipantViewResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetParticipantViewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetParticipantViewResponse struct{}"
	}

	return strings.Join([]string{"SetParticipantViewResponse", string(data)}, " ")
}
