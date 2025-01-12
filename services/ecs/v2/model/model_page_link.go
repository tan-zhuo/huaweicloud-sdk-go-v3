package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageLink
type PageLink struct {

	// 相应资源的链接。
	Href string `json:"href"`

	// 对应快捷链接。
	Rel string `json:"rel"`
}

func (o PageLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageLink struct{}"
	}

	return strings.Join([]string{"PageLink", string(data)}, " ")
}
