package model

import (
	"encoding/json"

	"strings"
)

// 搜索结果详情。
type SearchPictureItem struct {
	// 被搜索图片的路径。

	Path *string `json:"path,omitempty"`
	// 查询图片和被搜索图片的相似度，值越接近1表示越相似。

	Sim *float32 `json:"sim,omitempty"`
	// 自定义的标签名称和标签内容。

	Tags *interface{} `json:"tags,omitempty"`
}

func (o SearchPictureItem) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SearchPictureItem struct{}"
	}

	return strings.Join([]string{"SearchPictureItem", string(data)}, " ")
}