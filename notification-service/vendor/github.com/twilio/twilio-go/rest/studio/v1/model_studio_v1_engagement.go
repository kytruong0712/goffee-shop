/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Studio
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// StudioV1Engagement struct for StudioV1Engagement
type StudioV1Engagement struct {
	// The unique string that we created to identify the Engagement resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Engagement resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Flow.
	FlowSid *string `json:"flow_sid,omitempty"`
	// The SID of the Contact.
	ContactSid *string `json:"contact_sid,omitempty"`
	// The phone number, SIP address or Client identifier that triggered this Engagement. Phone numbers are in E.164 format (+16175551212). SIP addresses are formatted as `name@company.com`. Client identifiers are formatted `client:name`.
	ContactChannelAddress *string `json:"contact_channel_address,omitempty"`
	// The current state of the execution flow. As your flow executes, we save the state in a flow context. Your widgets can access the data in the flow context as variables, either in configuration fields or in text areas as variable substitution.
	Context *interface{} `json:"context,omitempty"`
	Status  *string      `json:"status,omitempty"`
	// The date and time in GMT when the Engagement was created in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the Engagement was updated in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
	// The URLs of the Engagement's nested resources.
	Links *map[string]interface{} `json:"links,omitempty"`
}
