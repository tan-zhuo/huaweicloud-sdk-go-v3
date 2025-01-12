package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneShowMappingRequest Request Object
type KeystoneShowMappingRequest struct {

	// 待查询的映射ID。
	Id string `json:"id"`
}

func (o KeystoneShowMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowMappingRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowMappingRequest", string(data)}, " ")
}
