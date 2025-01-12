package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceBy3rdResponse Response Object
type CreateInstanceBy3rdResponse struct {
	Result *InstancesResponseInstancesVoResult `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceBy3rdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceBy3rdResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceBy3rdResponse", string(data)}, " ")
}
