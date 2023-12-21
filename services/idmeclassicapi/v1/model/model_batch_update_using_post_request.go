package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateUsingPostRequest Request Object
type BatchUpdateUsingPostRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	// 数据模型的英文名称。
	ModelName string `json:"modelName"`

	Body *RdmParamVoListPersistableModelUpdateDto `json:"body,omitempty"`
}

func (o BatchUpdateUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateUsingPostRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateUsingPostRequest", string(data)}, " ")
}
