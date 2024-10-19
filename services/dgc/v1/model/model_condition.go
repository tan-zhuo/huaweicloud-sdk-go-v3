package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Condition struct {

	// 本节点依赖的前一个节点名称
	PreNodeName string `json:"preNodeName"`

	// EL表达式，如果EL表达式的计算结果为true，则触发执行本节点
	Expression string `json:"expression"`
}

func (o Condition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Condition struct{}"
	}

	return strings.Join([]string{"Condition", string(data)}, " ")
}
