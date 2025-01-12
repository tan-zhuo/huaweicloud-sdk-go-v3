package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetectLiveFaceByFileResponse Response Object
type DetectLiveFaceByFileResponse struct {
	Result *LiveDetectFaceRespResult `json:"result,omitempty"`

	// 警告信息列表。 调用失败时无此字段。
	WarningList *[]WarningList `json:"warning-list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DetectLiveFaceByFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectLiveFaceByFileResponse struct{}"
	}

	return strings.Join([]string{"DetectLiveFaceByFileResponse", string(data)}, " ")
}
