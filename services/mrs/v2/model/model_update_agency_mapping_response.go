package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAgencyMappingResponse Response Object
type UpdateAgencyMappingResponse struct {

	// 更新映射请求操作结果，succeeded为操作成功，failed为操作失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAgencyMappingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAgencyMappingResponse struct{}"
	}

	return strings.Join([]string{"UpdateAgencyMappingResponse", string(data)}, " ")
}
