/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Proxy
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ProxyV1ServicePhoneNumberCapabilities The capabilities of the phone number.
type ProxyV1ServicePhoneNumberCapabilities struct {
	Fax   bool `json:"fax,omitempty"`
	Mms   bool `json:"mms,omitempty"`
	Sms   bool `json:"sms,omitempty"`
	Voice bool `json:"voice,omitempty"`
}
