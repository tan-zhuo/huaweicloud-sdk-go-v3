package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetSpeedValueReq struct {

	// 限制带宽速率，单位 Kbps，-1表示不限速,1Mbps=1024Kbps。正整数表示限制到当前速率，电信支持限制速率:1Kbps,64 Kbps,256 Kbps,512Kbps,1Mbs，3Mbs,5Mbs,7Mbs,10Mbs,20Mbs,30Mbs,40Mbs,50Mbs,60Mbs,70Mbs,80Mbs,90Mbs,100Mbs,110Mbs,120Mbs,130Mbs,140Mbs,150Mbs。联通支持限制速率:256Kbps,512Kbps,1Mbps,2Mbps,7.25Mbps。
	SpeedValue int32 `json:"speed_value"`

	// iccid，传入的sim_card_id为0,则根据iccid进行处理
	Iccid *string `json:"iccid,omitempty"`
}

func (o SetSpeedValueReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSpeedValueReq struct{}"
	}

	return strings.Join([]string{"SetSpeedValueReq", string(data)}, " ")
}
