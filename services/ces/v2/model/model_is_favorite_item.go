package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IsFavoriteItem struct {

	// 监控看板是否标记收藏, true: 收藏, false: 未收藏
	IsFavorite *bool `json:"is_favorite,omitempty"`
}

func (o IsFavoriteItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IsFavoriteItem struct{}"
	}

	return strings.Join([]string{"IsFavoriteItem", string(data)}, " ")
}
