package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDictionaryRequest Request Object
type DeleteDictionaryRequest struct {

	// 字典ID
	DictId string `json:"dict_id"`

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o DeleteDictionaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDictionaryRequest struct{}"
	}

	return strings.Join([]string{"DeleteDictionaryRequest", string(data)}, " ")
}
