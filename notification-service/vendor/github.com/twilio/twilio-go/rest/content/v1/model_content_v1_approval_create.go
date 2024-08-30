/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Content
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ContentV1ApprovalCreate struct for ContentV1ApprovalCreate
type ContentV1ApprovalCreate struct {
	Name                *string `json:"name,omitempty"`
	Category            *string `json:"category,omitempty"`
	ContentType         *string `json:"content_type,omitempty"`
	Status              *string `json:"status,omitempty"`
	RejectionReason     *string `json:"rejection_reason,omitempty"`
	AllowCategoryChange *bool   `json:"allow_category_change,omitempty"`
}
