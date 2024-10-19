package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateGlobalEipRequestBodyGlobalEip 批量创建全域弹性公网IP请求体对象
type BatchCreateGlobalEipRequestBodyGlobalEip struct {

	// - 功能说明：全域弹性公网IP名称 - 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// - 功能说明：用户自定义的资源描述 - 约束：   - 值的长度最大512字符，由数字、字母、中文、_(下划线)、-（中划线）、.（点）组成。
	Description *string `json:"description,omitempty"`

	// 全域弹性公网IP池子名称
	GeipPoolName string `json:"geip_pool_name"`

	// 接入点信息
	AccessSite string `json:"access_site"`

	InternetBandwidthInfo *BatchCreateGlobalEipRequestBodyGlobalEipInternetBandwidthInfo `json:"internet_bandwidth_info"`

	// 批创个数
	Count *int32 `json:"count,omitempty"`

	// 全域弹性公网IP标签
	Tags *[]CreateGlobalEipRequestBodyGlobalEipTags `json:"tags,omitempty"`

	// - 企业项目ID。最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。 - 创建全域弹性公网IP时，给全域弹性公网IP绑定企业项目ID。 - 不指定该参数时，默认值是 0 - 关于企业项目ID的获取及企业项目特性的详细信息，请参见[《企业管理用户指南》](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0126101490.html)。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o BatchCreateGlobalEipRequestBodyGlobalEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGlobalEipRequestBodyGlobalEip struct{}"
	}

	return strings.Join([]string{"BatchCreateGlobalEipRequestBodyGlobalEip", string(data)}, " ")
}
