package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PostCustomTtsReq
type PostCustomTtsReq struct {

	// 待合成的文本，文本长度限制小于500字符。
	Text string `json:"text"`

	Config *TtsConfig `json:"config,omitempty"`
}

func (o PostCustomTtsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostCustomTtsReq struct{}"
	}

	return strings.Join([]string{"PostCustomTtsReq", string(data)}, " ")
}
