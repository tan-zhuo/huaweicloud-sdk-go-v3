/*
 * BCS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 通道信息
type ChannelCreateInfo struct {
	// 通道名称
	ChannelName string `json:"channel_name"`
	// 通道描述
	ChannelDescription *string `json:"channel_description,omitempty"`
}

func (o ChannelCreateInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ChannelCreateInfo", string(data)}, " ")
}