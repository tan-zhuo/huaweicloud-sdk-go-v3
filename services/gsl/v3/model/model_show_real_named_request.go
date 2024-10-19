package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRealNamedRequest Request Object
type ShowRealNamedRequest struct {

	// SIM卡标识，如果SIM卡标识传0则表示需要根据iccid处理。可通过[查询SIM卡列表接口](https://support.huaweicloud.com/api-ocgsl/gsl_07_0008.html)获取
	SimCardId int64 `json:"sim_card_id"`

	// iccid，传入的sim_card_id为0,则根据iccid进行处理
	Iccid *string `json:"iccid,omitempty"`
}

func (o ShowRealNamedRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRealNamedRequest struct{}"
	}

	return strings.Join([]string{"ShowRealNamedRequest", string(data)}, " ")
}
