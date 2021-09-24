package model

import (
	"encoding/json"

	"strings"
)

//
type Match struct {
	// 键。当前仅限定为resource_name

	Key string `json:"key"`
	// 值。每个值最大长度255个unicode字符，不能包含$ - . /等特殊字符。

	Value string `json:"value"`
}

func (o Match) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Match struct{}"
	}

	return strings.Join([]string{"Match", string(data)}, " ")
}