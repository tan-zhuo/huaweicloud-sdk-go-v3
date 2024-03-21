package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type BatchCreateInternetBandwidth struct {

	// 全域公网带宽的ID
	Id *string `json:"id,omitempty"`

	// - 功能说明：全域公网带宽名称 - 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// 全域弹性公网IP所属线路
	Isp string `json:"isp"`

	// 全域公网带宽大小（入云方向）
	IngressSize *int32 `json:"ingress_size,omitempty"`

	// 接入点信息
	AccessSite string `json:"access_site"`

	// 全域公网带宽大小（出云方向）
	Size int32 `json:"size"`

	// - 功能说明：用户自定义的资源描述 - 约束：   - 值的长度最大512字符，由数字、字母、中文、_(下划线)、-（中划线）、.（点）组成。
	Description *string `json:"description,omitempty"`

	// 计费模式
	ChargeMode *string `json:"charge_mode,omitempty"`

	// - 租户账号ID，获取租户账号ID请参见[租户账号ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)
	DomainId *string `json:"domain_id,omitempty"`

	// 状态
	Status *BatchCreateInternetBandwidthStatus `json:"status,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 是否创建成功标识，取值：successful、failed。
	RetStatus *string `json:"ret_status,omitempty"`

	// 全域公网带宽标签
	Tags *[]Tag `json:"tags,omitempty"`

	// 系统标签
	SysTags *[]Tag `json:"sys_tags,omitempty"`

	// - 企业项目ID。最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。 - 创建全域弹性公网IP时，给全域弹性公网IP绑定企业项目ID。 - 不指定该参数时，默认值是 0 - 关于企业项目ID的获取及企业项目特性的详细信息，请参见[《企业管理用户指南》](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0126101490.html)。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 全域公网带宽类型
	Type *string `json:"type,omitempty"`

	// 全域公网带宽资源的锁状态
	LockInfos *[]LockInfo `json:"lock_infos,omitempty"`
}

func (o BatchCreateInternetBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateInternetBandwidth struct{}"
	}

	return strings.Join([]string{"BatchCreateInternetBandwidth", string(data)}, " ")
}

type BatchCreateInternetBandwidthStatus struct {
	value string
}

type BatchCreateInternetBandwidthStatusEnum struct {
	NORMAL  BatchCreateInternetBandwidthStatus
	FREEZED BatchCreateInternetBandwidthStatus
}

func GetBatchCreateInternetBandwidthStatusEnum() BatchCreateInternetBandwidthStatusEnum {
	return BatchCreateInternetBandwidthStatusEnum{
		NORMAL: BatchCreateInternetBandwidthStatus{
			value: "NORMAL",
		},
		FREEZED: BatchCreateInternetBandwidthStatus{
			value: "FREEZED",
		},
	}
}

func (c BatchCreateInternetBandwidthStatus) Value() string {
	return c.value
}

func (c BatchCreateInternetBandwidthStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateInternetBandwidthStatus) UnmarshalJSON(b []byte) error {
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
