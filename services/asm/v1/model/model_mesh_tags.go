package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MeshTags 网格标签集合
type MeshTags struct {

	// Key值。  不能为空，最多支持128个字符  可用UTF-8格式表示的汉字、字母、数字和空格  支持部分特殊字符：_.:/=+-@  不能以\"_sys_\"开头
	Key *string `json:"key,omitempty"`

	// Value值。  可以为空但不能缺省，最多支持255个字符  可用UTF-8格式表示的汉字、字母、数字和空格  支持部分特殊字符：_.:/=+-@
	Value *string `json:"value,omitempty"`
}

func (o MeshTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MeshTags struct{}"
	}

	return strings.Join([]string{"MeshTags", string(data)}, " ")
}
