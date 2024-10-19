package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteArtifactFileResponse Response Object
type DeleteArtifactFileResponse struct {

	// 结果状态
	Status *string `json:"status,omitempty"`

	// 请求id
	TraceId *string `json:"trace_id,omitempty"`

	// 请求返回结果，接口不同，返回不同
	Result         *interface{} `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteArtifactFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteArtifactFileResponse struct{}"
	}

	return strings.Join([]string{"DeleteArtifactFileResponse", string(data)}, " ")
}
