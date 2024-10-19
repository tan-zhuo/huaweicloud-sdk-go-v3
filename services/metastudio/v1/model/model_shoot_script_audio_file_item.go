package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShootScriptAudioFileItem struct {

	// 剧本序号。
	SequenceNo int32 `json:"sequence_no"`

	// 语音驱动音频文件上传URL。创建和更新脚本时返回。单个文件最大100M。支持上传MP3/WAV/M4A文件。
	AudioFileUploadUrl *string `json:"audio_file_upload_url,omitempty"`

	// 语音驱动音频文件下载URL。查询脚本详情时返回。
	AudioFileDownloadUrl *string `json:"audio_file_download_url,omitempty"`
}

func (o ShootScriptAudioFileItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShootScriptAudioFileItem struct{}"
	}

	return strings.Join([]string{"ShootScriptAudioFileItem", string(data)}, " ")
}
