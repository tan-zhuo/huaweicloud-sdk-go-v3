package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelAssetTranscodeTaskResponse Response Object
type CancelAssetTranscodeTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelAssetTranscodeTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelAssetTranscodeTaskResponse struct{}"
	}

	return strings.Join([]string{"CancelAssetTranscodeTaskResponse", string(data)}, " ")
}
