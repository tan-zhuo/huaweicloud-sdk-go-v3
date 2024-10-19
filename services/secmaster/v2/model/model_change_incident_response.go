package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeIncidentResponse Response Object
type ChangeIncidentResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *IncidentDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeIncidentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeIncidentResponse struct{}"
	}

	return strings.Join([]string{"ChangeIncidentResponse", string(data)}, " ")
}
