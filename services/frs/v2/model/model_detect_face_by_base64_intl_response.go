package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DetectFaceByBase64IntlResponse Response Object
type DetectFaceByBase64IntlResponse struct {

	// 检测到的人脸。 调用失败时无此字段。
	Faces *[]DetectFace `json:"faces,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DetectFaceByBase64IntlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetectFaceByBase64IntlResponse struct{}"
	}

	return strings.Join([]string{"DetectFaceByBase64IntlResponse", string(data)}, " ")
}
