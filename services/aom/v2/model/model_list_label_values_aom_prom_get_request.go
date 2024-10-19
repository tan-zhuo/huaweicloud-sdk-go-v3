package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLabelValuesAomPromGetRequest Request Object
type ListLabelValuesAomPromGetRequest struct {

	// 查询所用标签。
	LabelName string `json:"label_name"`
}

func (o ListLabelValuesAomPromGetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLabelValuesAomPromGetRequest struct{}"
	}

	return strings.Join([]string{"ListLabelValuesAomPromGetRequest", string(data)}, " ")
}
