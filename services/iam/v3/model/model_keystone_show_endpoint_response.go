package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneShowEndpointResponse Response Object
type KeystoneShowEndpointResponse struct {
	Endpoint       *Endpoint `json:"endpoint,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o KeystoneShowEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowEndpointResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowEndpointResponse", string(data)}, " ")
}
