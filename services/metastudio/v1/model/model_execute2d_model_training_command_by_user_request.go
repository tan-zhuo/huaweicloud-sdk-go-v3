package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Execute2dModelTrainingCommandByUserRequest Request Object
type Execute2dModelTrainingCommandByUserRequest struct {

	// 任务ID。
	JobId string `json:"job_id"`

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	Body *Execute2dModelTrainingCommandByUserReq `json:"body,omitempty"`
}

func (o Execute2dModelTrainingCommandByUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Execute2dModelTrainingCommandByUserRequest struct{}"
	}

	return strings.Join([]string{"Execute2dModelTrainingCommandByUserRequest", string(data)}, " ")
}
