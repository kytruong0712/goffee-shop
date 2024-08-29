package authenticated

import (
	"net/http"

	userCtrl "github.com/kytruong0712/goffee-shop/api-gateway/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/httpserver"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

var (
	webErrIamIDIsRequired       = &httpserver.Error{Status: http.StatusBadRequest, Code: "iamid_is_required", Desc: "IamID is required"}
	webErrPhoneNumberIsRequired = &httpserver.Error{Status: http.StatusBadRequest, Code: "phone_number_is_required", Desc: "phone number is required"}
	webErrOTPIsRequired         = &httpserver.Error{Status: http.StatusBadRequest, Code: "otp_is_required", Desc: "otp is required"}
	webErrCountryCodeIsRequired = &httpserver.Error{Status: http.StatusBadRequest, Code: "country_code_is_required", Desc: "country code is required"}
	webErrInvalidCountryCode    = &httpserver.Error{Status: http.StatusBadRequest, Code: "invalid_country_code", Desc: "invalid country code"}
	webErrInvalidPhoneNumber    = &httpserver.Error{Status: http.StatusBadRequest, Code: "invalid_number", Desc: "Invalid phone number"}
)

var (
	webErrUserNotFound                    = &httpserver.Error{Status: http.StatusBadRequest, Code: "user_not_found", Desc: "User not found"}
	webErrIncorrectRegisteredPhoneNumber  = &httpserver.Error{Status: http.StatusBadRequest, Code: "incorrect_registered_phone_number", Desc: "Incorrect registered phone number"}
	webErrInvalidUpdateProfileRequestData = &httpserver.Error{Status: http.StatusBadRequest, Code: "invalid_update_profile_request_data", Desc: "Invalid update profile request data"}
	webErrUserAlreadyActivated            = &httpserver.Error{Status: http.StatusBadRequest, Code: "user_already_activated", Desc: "User already activated"}
	webErrOTPIsNotMatched                 = &httpserver.Error{Status: http.StatusBadRequest, Code: "otp_is_not_matched", Desc: "OTP is not matched"}
	webErrOTPIsExpired                    = &httpserver.Error{Status: http.StatusBadRequest, Code: "otp_is_expired", Desc: "OTP is expired"}
)

func convertCtrlErr(err error) error {
	// gRPC errors
	st, ok := status.FromError(err)
	if ok {
		for _, detail := range st.Details() {
			switch t := detail.(type) {
			case *errdetails.BadRequest:
				for _, violation := range t.FieldViolations {
					switch violation.Description {
					case userCtrl.ErrIamIDIsRequired.Error():
						return webErrIamIDIsRequired
					case userCtrl.ErrPhoneNumberAlreadyExists.Error():
						return webErrPhoneNumberIsRequired
					case userCtrl.ErrUserNotFound.Error():
						return webErrUserNotFound
					case userCtrl.ErrUserAlreadyActivated.Error():
						return webErrUserAlreadyActivated
					case userCtrl.ErrOTPIsNotMatched.Error():
						return webErrOTPIsNotMatched
					case userCtrl.ErrOTPIsExpired.Error():
						return webErrOTPIsExpired
					}
				}
			}
		}
	}

	// Non-gRPC errors
	switch err.(type) {
	// TODO: handle non-gRPC error
	}

	return &httpserver.Error{Status: http.StatusInternalServerError, Code: "internal_error", Desc: err.Error()}
}
