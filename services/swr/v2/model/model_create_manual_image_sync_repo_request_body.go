package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateManualImageSyncRepoRequestBody 手动同步镜像需要的参数
type CreateManualImageSyncRepoRequestBody struct {

	// 版本列表
	ImageTag []string `json:"imageTag"`

	// 是否覆盖，默认为false
	Override *bool `json:"override,omitempty"`

	// 目标组织
	RemoteNamespace string `json:"remoteNamespace"`

	// 目标region ID。
	RemoteRegionId string `json:"remoteRegionId"`
}

func (o CreateManualImageSyncRepoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManualImageSyncRepoRequestBody struct{}"
	}

	return strings.Join([]string{"CreateManualImageSyncRepoRequestBody", string(data)}, " ")
}
