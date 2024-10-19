package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type OpenMysqlProxyRequestBody struct {

	// 代理规格码。
	FlavorRef string `json:"flavor_ref"`

	// 代理实例节点数，取值整数2-32。
	NodeNum int32 `json:"node_num"`

	// 代理实例名称。用于表示实例的名称，同一租户下，同类型的实例名可重名。  取值范围：最小为4个字符，最大为64个字符且不超过64个字节（注意：一个中文字符占用3个字节），必须以字母或中文开头，区分大小写，可以包含字母、数字、中划线、下划线或中文，不能包含其他特殊字符。
	ProxyName *string `json:"proxy_name,omitempty"`

	// 代理实例类型。默认类型为readwrite。
	ProxyMode *OpenMysqlProxyRequestBodyProxyMode `json:"proxy_mode,omitempty"`

	// 数据库代理路由模式，默认为权重负载模式。  取值范围: - 0，表示权重负载模式; - 1，表示负载均衡模式（数据库主节点不接受读请求）； - 2，表示负载均衡模式（数据库主节点接受读请求）。
	RouteMode *int32 `json:"route_mode,omitempty"`

	// 数据库节点的读权重设置。  在proxy_mode为readonly时，只能为只读节点选择权重。
	NodesReadWeight *[]NodesWeight `json:"nodes_read_weight,omitempty"`

	// 数据库VPC下的子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 是否开启新增节点自动加入该Proxy。如果需要设置是否开启新增节点自动加入该Proxy，请联系客服人员添加白名单，加入白名单后，方可输入该字段。  取值范围： - ON：开启。 - OFF：关闭。
	NewNodeAutoAddStatus *string `json:"new_node_auto_add_status,omitempty"`

	// 新增节点的读权重：    - 如果路由模式为0，新增节点自动加入为ON，取值为0~1000。 - 如果路由模式不为0或新增节点自动加入为OFF，则可不输入读权重。
	NewNodeWeight *int32 `json:"new_node_weight,omitempty"`
}

func (o OpenMysqlProxyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenMysqlProxyRequestBody struct{}"
	}

	return strings.Join([]string{"OpenMysqlProxyRequestBody", string(data)}, " ")
}

type OpenMysqlProxyRequestBodyProxyMode struct {
	value string
}

type OpenMysqlProxyRequestBodyProxyModeEnum struct {
	READWRITE OpenMysqlProxyRequestBodyProxyMode
	READONLY  OpenMysqlProxyRequestBodyProxyMode
}

func GetOpenMysqlProxyRequestBodyProxyModeEnum() OpenMysqlProxyRequestBodyProxyModeEnum {
	return OpenMysqlProxyRequestBodyProxyModeEnum{
		READWRITE: OpenMysqlProxyRequestBodyProxyMode{
			value: "readwrite",
		},
		READONLY: OpenMysqlProxyRequestBodyProxyMode{
			value: "readonly",
		},
	}
}

func (c OpenMysqlProxyRequestBodyProxyMode) Value() string {
	return c.value
}

func (c OpenMysqlProxyRequestBodyProxyMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenMysqlProxyRequestBodyProxyMode) UnmarshalJSON(b []byte) error {
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
