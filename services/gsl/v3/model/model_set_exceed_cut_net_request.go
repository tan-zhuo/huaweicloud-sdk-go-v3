package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetExceedCutNetRequest Request Object
type SetExceedCutNetRequest struct {

	// SIM卡标识，如果SIM卡标识传0则表示需要根据iccid处理。可通过[查询SIM卡列表接口](https://support.huaweicloud.com/api-ocgsl/gsl_07_0008.html)获取
	SimCardId int64 `json:"sim_card_id"`

	Body *ExceedCutNetReq `json:"body,omitempty"`
}

func (o SetExceedCutNetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetExceedCutNetRequest struct{}"
	}

	return strings.Join([]string{"SetExceedCutNetRequest", string(data)}, " ")
}
