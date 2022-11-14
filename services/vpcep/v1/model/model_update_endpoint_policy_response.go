package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateEndpointPolicyResponse struct {

	// 终端节点的ID，唯一标识。
	Id *string `json:"id,omitempty"`

	// 终端节点连接的终端节点服务类型。 ● gataway：由运维人员配置。用户无需创建，可直接使用。 ● interface：包括运维人员配置的云服务和用户自己创建的私有服务。 其中，运维人员配置的云服务无需创建，用户可直接使用。 您可以通过查询公共终端节点服务列表， 查看由运维人员配置的所有用户可见且可连接的终端节点服务， 并通过创建终端节点服务创建Interface类型的终端节点服务。
	ServiceType *UpdateEndpointPolicyResponseServiceType `json:"service_type,omitempty"`

	// 终端节点的连接状态。 ● pendingAcceptance：待接受 ● creating：创建中 ● accepted：已接受 ● failed：失败
	Status *UpdateEndpointPolicyResponseStatus `json:"status,omitempty"`

	// 帐号状态。 ● frozen：冻结 ● active：解冻
	ActiveStatus *[]string `json:"active_status,omitempty"`

	// 终端节点服务的名称。
	EndpointServiceName *string `json:"endpoint_service_name,omitempty"`

	// 终端节点的报文标识。
	MarkerId *int32 `json:"marker_id,omitempty"`

	// 终端节点服务的ID。
	EndpointServiceId *string `json:"endpoint_service_id,omitempty"`

	// 是否创建域名。 ● true：创建域名 ● false：不创建域名 说明 当创建连接gateway类型终端节点服务的终端节点时， “enable_dns”设置为true或者false，均不创建域名。
	EnableDns *bool `json:"enable_dns,omitempty"`

	// 访问所连接的终端节点服务的域名。 当“enable_dns”为true时，该参数可见。
	DnsNames *[]string `json:"dns_names,omitempty"`

	// 访问所连接的终端节点服务的IP。 仅当同时满足如下条件时，返回该参数： 当查询连接interface类型终端节点服务的终端节点时。 终端节点服务启用“连接审批”功能，且已经“接受”连接审批。 “status”可以是“accepted”或者“rejected（仅支持“接受”连接审批后再“拒绝”的情况）”。
	Ip *string `json:"ip,omitempty"`

	// 终端节点所在的VPC的ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// vpc_id对应VPC下已创建的网络（network）的ID，UUID格式。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 终端节点的创建时间。 采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ
	CreatedAt *string `json:"created_at,omitempty"`

	// 终端节点的更新时间。 采用UTC时间格式，格式为：YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 项目ID，获取方法请参见获取项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 标签列表，没有标签默认为空数组。
	Tags *[]TagList `json:"tags,omitempty"`

	// 错误信息。 当终端节点状态异常，即“status”的值为“failed”时，会返回该字段。
	Error *[]QueryError `json:"error,omitempty"`

	// 控制访问终端节点的白名单。 若未创建，则返回空列表。 创建连接Interface类型终端节点服务的终端节点时，显示此参数。
	Whitelist *[]string `json:"whitelist,omitempty"`

	// 是否开启网络ACL隔离。 ● true：开启网络ACL隔离 ● false：不开启网络ACL隔离 若未指定，则返回false。 创建连接Interface类型终端节点服务的终端节点时，显示此参数。
	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`

	// 路由表ID列表。 若未指定，返回默认VPC下路由表ID。 创建连接Gateway类型终端节点服务的终端节点时，显示此参数。
	Routetables *[]string `json:"routetables,omitempty"`

	// 描述字段，支持中英文字母、数字等字符，不支持“<”或“>”字符。
	Description *string `json:"description,omitempty"`

	// 只涉及开启双端固定的网关型终端节点，响应体展示此字段
	PolicyStatement *[]PolicyStatement `json:"policy_statement,omitempty"`

	// 终端节点相关联的Pood的ID
	EndpointPoolId *string `json:"endpoint_pool_id,omitempty"`

	// 终端节点关联的Public Border Group信息，只有当终端节点和边缘Pool相关联时才会返回改字段
	PublicBorderGroup *string `json:"public_border_group,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o UpdateEndpointPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointPolicyResponse", string(data)}, " ")
}

type UpdateEndpointPolicyResponseServiceType struct {
	value string
}

type UpdateEndpointPolicyResponseServiceTypeEnum struct {
	INTERFACE UpdateEndpointPolicyResponseServiceType
	GATEWAY   UpdateEndpointPolicyResponseServiceType
}

func GetUpdateEndpointPolicyResponseServiceTypeEnum() UpdateEndpointPolicyResponseServiceTypeEnum {
	return UpdateEndpointPolicyResponseServiceTypeEnum{
		INTERFACE: UpdateEndpointPolicyResponseServiceType{
			value: "interface",
		},
		GATEWAY: UpdateEndpointPolicyResponseServiceType{
			value: "gateway",
		},
	}
}

func (c UpdateEndpointPolicyResponseServiceType) Value() string {
	return c.value
}

func (c UpdateEndpointPolicyResponseServiceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEndpointPolicyResponseServiceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type UpdateEndpointPolicyResponseStatus struct {
	value string
}

type UpdateEndpointPolicyResponseStatusEnum struct {
	PENDING_ACCEPTANCE UpdateEndpointPolicyResponseStatus
	CREATING           UpdateEndpointPolicyResponseStatus
	ACCEPTED           UpdateEndpointPolicyResponseStatus
	REJECTED           UpdateEndpointPolicyResponseStatus
	FAILED             UpdateEndpointPolicyResponseStatus
	DELETING           UpdateEndpointPolicyResponseStatus
}

func GetUpdateEndpointPolicyResponseStatusEnum() UpdateEndpointPolicyResponseStatusEnum {
	return UpdateEndpointPolicyResponseStatusEnum{
		PENDING_ACCEPTANCE: UpdateEndpointPolicyResponseStatus{
			value: "pendingAcceptance",
		},
		CREATING: UpdateEndpointPolicyResponseStatus{
			value: "creating",
		},
		ACCEPTED: UpdateEndpointPolicyResponseStatus{
			value: "accepted",
		},
		REJECTED: UpdateEndpointPolicyResponseStatus{
			value: "rejected",
		},
		FAILED: UpdateEndpointPolicyResponseStatus{
			value: "failed",
		},
		DELETING: UpdateEndpointPolicyResponseStatus{
			value: "deleting",
		},
	}
}

func (c UpdateEndpointPolicyResponseStatus) Value() string {
	return c.value
}

func (c UpdateEndpointPolicyResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEndpointPolicyResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}