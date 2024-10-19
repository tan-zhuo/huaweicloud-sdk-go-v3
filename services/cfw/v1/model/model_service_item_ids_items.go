package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ServiceItemIdsItems struct {

	// id值
	Id *string `json:"id,omitempty"`
}

func (o ServiceItemIdsItems) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceItemIdsItems struct{}"
	}

	return strings.Join([]string{"ServiceItemIdsItems", string(data)}, " ")
}
