package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskHostingResult 输出为hosting时，hosting结果文件的相关信息
type TaskHostingResult struct {
	HostingResult *TaskHostingResultHostingResult `json:"hosting_result,omitempty"`
}

func (o TaskHostingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskHostingResult struct{}"
	}

	return strings.Join([]string{"TaskHostingResult", string(data)}, " ")
}
