package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneShowSecurityComplianceByOptionResponse Response Object
type KeystoneShowSecurityComplianceByOptionResponse struct {
	Config         *ConfigByOption `json:"config,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o KeystoneShowSecurityComplianceByOptionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowSecurityComplianceByOptionResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowSecurityComplianceByOptionResponse", string(data)}, " ")
}
