package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFileRequest Request Object
type CreateFileRequest struct {

	// 授权名称。
	XRepoAuth string `json:"X-Repo-Auth"`

	// 命名空间ID或者URL编码名称。
	Namespace string `json:"namespace"`

	// 仓库项目ID，如果含有“/”，需要将“/”替换为“:”。
	Project string `json:"project"`

	// 文件路径，需要将“/”替换为“:”。
	Path string `json:"path"`

	// 分支名称或者tag标签名称或者commit sha。
	Ref string `json:"ref"`

	Body *FileCreate `json:"body,omitempty"`
}

func (o CreateFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFileRequest struct{}"
	}

	return strings.Join([]string{"CreateFileRequest", string(data)}, " ")
}
