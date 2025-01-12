package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneShowServiceResponse Response Object
type KeystoneShowServiceResponse struct {
	Service        *Service `json:"service,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o KeystoneShowServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowServiceResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowServiceResponse", string(data)}, " ")
}
