package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompareFace struct {
	BoundingBox *BoundingBox `json:"bounding_box"`
}

func (o CompareFace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareFace struct{}"
	}

	return strings.Join([]string{"CompareFace", string(data)}, " ")
}
