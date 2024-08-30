package public

import (
	"net/http"

	userCtrl "github.com/kytruong0712/goffee-shop/api-gateway/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/httpserver"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

var (
	webErrFullNameIsRequired    = &httpserver.Error{Status: http.StatusBadRequest, Code: "full_name_is_required", Desc: "Full name is required"}
	webErrPhoneNumberIsRequired = &httpserver.Error{Status: http.StatusBadRequest, Code: "phone_number_is_required", Desc: "Phone number is required"}
	webErrInvalidPhoneNumber    = &httpserver.Error{Status: http.StatusBadRequest, Code: "invalid_number", Desc: "Invalid phone number"}
	webErrPasswordIsRequired    = &httpserver.Error{Status: http.StatusBadRequest, Code: "password_is_required", Desc: "Password is required"}
	webErrInvalidPassword       = &httpserver.Error{Status: http.StatusBadRequest, Code: "invalid_password", Desc: "Invalid password. Rules: The password should between 8 and 12 characters, contains at least one uppercase letter, one lowercase letter, one number and one special character"}
)

var (
	webErrPhoneNumberAlreadyExists     = &httpserver.Error{Status: http.StatusBadRequest, Code: "phone_number_already_exists", Desc: "Phone number already exists"}
	webErrLoginIDOrPasswordIsIncorrect = &httpserver.Error{Status: http.StatusBadRequest, Code: "account_name_or_password_is_incorrect", Desc: "Account name or password is incorrect"}
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
					case userCtrl.ErrLoginIDOrPasswordIsIncorrect.Error():
						return webErrLoginIDOrPasswordIsIncorrect
					case userCtrl.ErrPhoneNumberAlreadyExists.Error():
						return webErrPhoneNumberAlreadyExists
					case userCtrl.ErrFullNameIsRequired.Error():
						return webErrFullNameIsRequired
					case userCtrl.ErrPhoneNumberIsRequired.Error():
						return webErrPhoneNumberIsRequired
					case userCtrl.ErrInvalidPhoneNumber.Error():
						return webErrInvalidPhoneNumber
					case userCtrl.ErrPasswordIsRequired.Error():
						return webErrPasswordIsRequired
					case userCtrl.ErrInvalidPassword.Error():
						return webErrInvalidPassword
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
