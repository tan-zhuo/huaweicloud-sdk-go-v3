package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneShowVersionRequest Request Object
type KeystoneShowVersionRequest struct {
}

func (o KeystoneShowVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowVersionRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowVersionRequest", string(data)}, " ")
}
