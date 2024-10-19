package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBizMetricResponse Response Object
type CreateBizMetricResponse struct {
	Data           *CreateBizMetricResultData `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o CreateBizMetricResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBizMetricResponse struct{}"
	}

	return strings.Join([]string{"CreateBizMetricResponse", string(data)}, " ")
}
