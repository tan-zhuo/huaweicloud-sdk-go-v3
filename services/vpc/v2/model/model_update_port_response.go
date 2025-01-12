package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePortResponse Response Object
type UpdatePortResponse struct {
	Port           *Port `json:"port,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdatePortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortResponse struct{}"
	}

	return strings.Join([]string{"UpdatePortResponse", string(data)}, " ")
}
