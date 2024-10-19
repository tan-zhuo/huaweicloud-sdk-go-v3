package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FileRedirectionOptionsCompressionSwitchOptions 压缩开关控制项。
type FileRedirectionOptionsCompressionSwitchOptions struct {

	// 压缩阈值（Byte）。取值范围为[0-10240]。默认：512。
	CompressionThreshold *int32 `json:"compression_threshold,omitempty"`

	// 最小压缩率。取值范围为[0-1000]。默认：900。
	MinimumCompressionRate *int32 `json:"minimum_compression_rate,omitempty"`
}

func (o FileRedirectionOptionsCompressionSwitchOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FileRedirectionOptionsCompressionSwitchOptions struct{}"
	}

	return strings.Join([]string{"FileRedirectionOptionsCompressionSwitchOptions", string(data)}, " ")
}
