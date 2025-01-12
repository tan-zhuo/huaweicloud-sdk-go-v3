package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQualityEnhanceTemplateResponse Response Object
type CreateQualityEnhanceTemplateResponse struct {

	// 模板ID。
	TemplateId     *int32 `json:"template_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateQualityEnhanceTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQualityEnhanceTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateQualityEnhanceTemplateResponse", string(data)}, " ")
}
