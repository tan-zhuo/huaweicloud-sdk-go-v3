package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteQualityEnhanceTemplateResponse Response Object
type DeleteQualityEnhanceTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteQualityEnhanceTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQualityEnhanceTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteQualityEnhanceTemplateResponse", string(data)}, " ")
}
