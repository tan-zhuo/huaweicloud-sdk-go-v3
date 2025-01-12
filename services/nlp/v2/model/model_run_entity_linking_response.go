package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunEntityLinkingResponse Response Object
type RunEntityLinkingResponse struct {

	// 识别出的实体列表
	Entities       *[]LinkedEntity `json:"entities,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o RunEntityLinkingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunEntityLinkingResponse struct{}"
	}

	return strings.Join([]string{"RunEntityLinkingResponse", string(data)}, " ")
}
