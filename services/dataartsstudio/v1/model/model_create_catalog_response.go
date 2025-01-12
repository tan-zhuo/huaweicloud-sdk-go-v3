package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCatalogResponse Response Object
type CreateCatalogResponse struct {
	Data           *CreateCatalogResultData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CreateCatalogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCatalogResponse struct{}"
	}

	return strings.Join([]string{"CreateCatalogResponse", string(data)}, " ")
}
