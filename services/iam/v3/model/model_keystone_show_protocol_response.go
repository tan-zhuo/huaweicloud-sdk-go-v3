package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneShowProtocolResponse Response Object
type KeystoneShowProtocolResponse struct {
	Protocol       *ProtocolResult `json:"protocol,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o KeystoneShowProtocolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowProtocolResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowProtocolResponse", string(data)}, " ")
}
