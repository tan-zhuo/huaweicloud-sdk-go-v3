package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAppGroupsRequestBody struct {

	// 分组名称
	Name string `json:"name"`
}

func (o UpdateAppGroupsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppGroupsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAppGroupsRequestBody", string(data)}, " ")
}
