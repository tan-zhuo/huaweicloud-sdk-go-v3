package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVariableRequest Request Object
type UpdateVariableRequest struct {

	// 测试工程id
	TestSuiteId int32 `json:"test_suite_id"`

	Body *[]UpdateVariableRequestBody `json:"body,omitempty"`
}

func (o UpdateVariableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVariableRequest struct{}"
	}

	return strings.Join([]string{"UpdateVariableRequest", string(data)}, " ")
}
