package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeTransportationLicenseResponse struct {
	Result         *TransportationLicenseResult `json:"result,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o RecognizeTransportationLicenseResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeTransportationLicenseResponse struct{}"
	}

	return strings.Join([]string{"RecognizeTransportationLicenseResponse", string(data)}, " ")
}