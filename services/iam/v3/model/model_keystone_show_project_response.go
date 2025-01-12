package model

import (
	"github.com/tan-zhuo/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneShowProjectResponse Response Object
type KeystoneShowProjectResponse struct {
	Project        *ProjectResult `json:"project,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o KeystoneShowProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowProjectResponse struct{}"
	}

	return strings.Join([]string{"KeystoneShowProjectResponse", string(data)}, " ")
}
