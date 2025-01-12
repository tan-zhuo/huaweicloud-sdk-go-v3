package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyAssumedby
type AgencyAssumedby struct {
	User *AgencyAssumedbyUser `json:"user"`
}

func (o AgencyAssumedby) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyAssumedby struct{}"
	}

	return strings.Join([]string{"AgencyAssumedby", string(data)}, " ")
}
