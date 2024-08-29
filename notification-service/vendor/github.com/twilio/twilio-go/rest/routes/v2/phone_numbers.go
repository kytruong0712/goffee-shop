/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Routes
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
	"strings"
)

// Fetch the Inbound Processing Region assigned to a phone number.
func (c *ApiService) FetchPhoneNumber(PhoneNumber string) (*RoutesV2PhoneNumber, error) {
	path := "/v2/PhoneNumbers/{PhoneNumber}"
	path = strings.Replace(path, "{"+"PhoneNumber"+"}", PhoneNumber, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &RoutesV2PhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdatePhoneNumber'
type UpdatePhoneNumberParams struct {
	// The Inbound Processing Region used for this phone number for voice
	VoiceRegion *string `json:"VoiceRegion,omitempty"`
	// A human readable description of this resource, up to 64 characters.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdatePhoneNumberParams) SetVoiceRegion(VoiceRegion string) *UpdatePhoneNumberParams {
	params.VoiceRegion = &VoiceRegion
	return params
}
func (params *UpdatePhoneNumberParams) SetFriendlyName(FriendlyName string) *UpdatePhoneNumberParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Assign an Inbound Processing Region to a phone number.
func (c *ApiService) UpdatePhoneNumber(PhoneNumber string, params *UpdatePhoneNumberParams) (*RoutesV2PhoneNumber, error) {
	path := "/v2/PhoneNumbers/{PhoneNumber}"
	path = strings.Replace(path, "{"+"PhoneNumber"+"}", PhoneNumber, -1)

	data := url.Values{}
	headers := map[string]interface{}{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	if params != nil && params.VoiceRegion != nil {
		data.Set("VoiceRegion", *params.VoiceRegion)
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &RoutesV2PhoneNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
