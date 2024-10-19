package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetScriptJobBatchResponse Response Object
type GetScriptJobBatchResponse struct {
	Data           *JobScriptBatchDetailModel `json:"data,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o GetScriptJobBatchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetScriptJobBatchResponse struct{}"
	}

	return strings.Join([]string{"GetScriptJobBatchResponse", string(data)}, " ")
}
