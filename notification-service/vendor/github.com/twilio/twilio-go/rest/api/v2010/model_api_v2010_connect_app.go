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

// ApiV2010ConnectApp struct for ApiV2010ConnectApp
type ApiV2010ConnectApp struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the ConnectApp resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The URL we redirect the user to after we authenticate the user and obtain authorization to access the Connect App.
	AuthorizeRedirectUrl *string `json:"authorize_redirect_url,omitempty"`
	// The company name set for the Connect App.
	CompanyName *string `json:"company_name,omitempty"`
	// The HTTP method we use to call `deauthorize_callback_url`.
	DeauthorizeCallbackMethod *string `json:"deauthorize_callback_method,omitempty"`
	// The URL we call using the `deauthorize_callback_method` to de-authorize the Connect App.
	DeauthorizeCallbackUrl *string `json:"deauthorize_callback_url,omitempty"`
	// The description of the Connect App.
	Description *string `json:"description,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The public URL where users can obtain more information about this Connect App.
	HomepageUrl *string `json:"homepage_url,omitempty"`
	// The set of permissions that your ConnectApp requests.
	Permissions *[]string `json:"permissions,omitempty"`
	// The unique string that that we created to identify the ConnectApp resource.
	Sid *string `json:"sid,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`.
	Uri *string `json:"uri,omitempty"`
}
