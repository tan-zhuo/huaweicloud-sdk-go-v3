package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneShowSecurityComplianceResponse Response Object
type KeystoneShowSecurityComplianceResponse struct {
	Config         *Config `json:"config,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o KeystoneShowSecurityComplianceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowSecurityComplianceResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowSecurityComplianceResponse", string(data)}, " ")
}
