package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGaussMySqlJobInfoResponse Response Object
type ShowGaussMySqlJobInfoResponse struct {
	Job            *GetJobInfoDetail `json:"job,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowGaussMySqlJobInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlJobInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlJobInfoResponse", string(data)}, " ")
}
