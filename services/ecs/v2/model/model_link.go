package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Link 相关快捷链接地址。
type Link struct {

	// 对应快捷链接。
	Href string `json:"href"`

	// 快捷链接标记名称。
	Rel string `json:"rel"`
}

func (o Link) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Link struct{}"
	}

	return strings.Join([]string{"Link", string(data)}, " ")
}
