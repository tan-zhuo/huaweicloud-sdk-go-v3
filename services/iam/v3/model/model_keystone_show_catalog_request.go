package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneShowCatalogRequest Request Object
type KeystoneShowCatalogRequest struct {
}

func (o KeystoneShowCatalogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowCatalogRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowCatalogRequest", string(data)}, " ")
}
