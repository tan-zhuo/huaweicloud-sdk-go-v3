package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryRequest Request Object
type ListHistoryRequest struct {

	// 存储ID
	DataStoreId string `json:"data_store_id"`

	Body *GetHistoryRequest `json:"body,omitempty"`
}

func (o ListHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryRequest struct{}"
	}

	return strings.Join([]string{"ListHistoryRequest", string(data)}, " ")
}
