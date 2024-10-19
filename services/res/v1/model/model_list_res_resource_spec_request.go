package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResResourceSpecRequest Request Object
type ListResResourceSpecRequest struct {
}

func (o ListResResourceSpecRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResResourceSpecRequest struct{}"
	}

	return strings.Join([]string{"ListResResourceSpecRequest", string(data)}, " ")
}
