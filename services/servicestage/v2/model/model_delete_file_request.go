package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFileRequest Request Object
type DeleteFileRequest struct {

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

	// 提交描述。
	Message string `json:"message"`

	// 最后一次提交的commit sha值。
	Sha string `json:"sha"`
}

func (o DeleteFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFileRequest struct{}"
	}

	return strings.Join([]string{"DeleteFileRequest", string(data)}, " ")
}
