package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportImageQuickRequest Request Object
type ImportImageQuickRequest struct {
	Body *QuickImportImageByFileRequestBody `json:"body,omitempty"`
}

func (o ImportImageQuickRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportImageQuickRequest struct{}"
	}

	return strings.Join([]string{"ImportImageQuickRequest", string(data)}, " ")
}
