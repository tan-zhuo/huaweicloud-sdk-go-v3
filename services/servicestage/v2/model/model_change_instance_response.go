package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeInstanceResponse Response Object
type ChangeInstanceResponse struct {

	// Job ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceResponse struct{}"
	}

	return strings.Join([]string{"ChangeInstanceResponse", string(data)}, " ")
}
