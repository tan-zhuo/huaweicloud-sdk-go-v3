package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateReq This is a auto create Body Object
type UpdateReq struct {

	// 视频名，长度3~63位。 大小写字母，数字，汉字及部分符号(“_”、“-”、“#”)组成
	Name string `json:"name"`
}

func (o UpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateReq struct{}"
	}

	return strings.Join([]string{"UpdateReq", string(data)}, " ")
}
