package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunSqlConversionRequest Request Object
type RunSqlConversionRequest struct {
	Body *SqlConvertReq `json:"body,omitempty"`
}

func (o RunSqlConversionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSqlConversionRequest struct{}"
	}

	return strings.Join([]string{"RunSqlConversionRequest", string(data)}, " ")
}
