package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RecordRuleReq 修改录制规则
type RecordRuleReq struct {
	ObsAddr *RecordObsFileAddr `json:"obs_addr"`

	//  录制格式：支持HLS格式和MP4格式（HLS和MP4为大写）。   - 若配置HLS则必须携带HLSRecordConfig参数  - 若配置MP4则需要携带MP4RecordConfig
	RecordFormats []RecordRuleReqRecordFormats `json:"record_formats"`

	HlsConfig *HlsRecordConfig `json:"hls_config,omitempty"`

	Mp4Config *Mp4RecordConfig `json:"mp4_config,omitempty"`
}

func (o RecordRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordRuleReq struct{}"
	}

	return strings.Join([]string{"RecordRuleReq", string(data)}, " ")
}

type RecordRuleReqRecordFormats struct {
	value string
}

type RecordRuleReqRecordFormatsEnum struct {
	HLS RecordRuleReqRecordFormats
	MP4 RecordRuleReqRecordFormats
}

func GetRecordRuleReqRecordFormatsEnum() RecordRuleReqRecordFormatsEnum {
	return RecordRuleReqRecordFormatsEnum{
		HLS: RecordRuleReqRecordFormats{
			value: "HLS",
		},
		MP4: RecordRuleReqRecordFormats{
			value: "MP4",
		},
	}
}

func (c RecordRuleReqRecordFormats) Value() string {
	return c.value
}

func (c RecordRuleReqRecordFormats) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecordRuleReqRecordFormats) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
