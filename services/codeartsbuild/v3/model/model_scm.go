package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Scm 代码信息
type Scm struct {

	// 代码Tag
	BuildTag *string `json:"build_tag,omitempty"`

	// 代码提交ID
	BuildCommitId *string `json:"build_commit_id,omitempty"`
}

func (o Scm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Scm struct{}"
	}

	return strings.Join([]string{"Scm", string(data)}, " ")
}
