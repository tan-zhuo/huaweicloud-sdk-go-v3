package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppFavorite 是否收藏应用，收藏的应用会在列表里优先显示
type AppFavorite struct {
}

func (o AppFavorite) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppFavorite struct{}"
	}

	return strings.Join([]string{"AppFavorite", string(data)}, " ")
}
