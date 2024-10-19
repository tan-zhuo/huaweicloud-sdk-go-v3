package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlanceShowImageSchemasRequest Request Object
type GlanceShowImageSchemasRequest struct {
}

func (o GlanceShowImageSchemasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageSchemasRequest struct{}"
	}

	return strings.Join([]string{"GlanceShowImageSchemasRequest", string(data)}, " ")
}
