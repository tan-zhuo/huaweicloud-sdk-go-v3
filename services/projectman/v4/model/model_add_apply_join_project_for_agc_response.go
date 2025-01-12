package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddApplyJoinProjectForAgcResponse Response Object
type AddApplyJoinProjectForAgcResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddApplyJoinProjectForAgcResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddApplyJoinProjectForAgcResponse struct{}"
	}

	return strings.Join([]string{"AddApplyJoinProjectForAgcResponse", string(data)}, " ")
}
