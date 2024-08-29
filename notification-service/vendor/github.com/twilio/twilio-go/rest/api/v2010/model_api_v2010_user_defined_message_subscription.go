/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ApiV2010UserDefinedMessageSubscription struct for ApiV2010UserDefinedMessageSubscription
type ApiV2010UserDefinedMessageSubscription struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that subscribed to the User Defined Messages.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the [Call](https://www.twilio.com/docs/voice/api/call-resource) the User Defined Message Subscription is associated with. This refers to the Call SID that is producing the User Defined Messages.
	CallSid *string `json:"call_sid,omitempty"`
	// The SID that uniquely identifies this User Defined Message Subscription.
	Sid *string `json:"sid,omitempty"`
	// The date that this User Defined Message Subscription was created, given in RFC 2822 format.
	DateCreated *string `json:"date_created,omitempty"`
	// The URI of the User Defined Message Subscription Resource, relative to `https://api.twilio.com`.
	Uri *string `json:"uri,omitempty"`
}
