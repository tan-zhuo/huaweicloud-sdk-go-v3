package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckProjectNameV4Request Request Object
type CheckProjectNameV4Request struct {
	Body *CheckProjectNameRequestV4 `json:"body,omitempty"`
}

func (o CheckProjectNameV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckProjectNameV4Request struct{}"
	}

	return strings.Join([]string{"CheckProjectNameV4Request", string(data)}, " ")
}
