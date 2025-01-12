package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneShowServiceRequest Request Object
type KeystoneShowServiceRequest struct {

	// 待查询的服务ID。
	ServiceId string `json:"service_id"`
}

func (o KeystoneShowServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowServiceRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowServiceRequest", string(data)}, " ")
}
