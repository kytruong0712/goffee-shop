package public

import (
	"net/http"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/services/user"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/httpserver"
)

var (
	webErrFullNameIsRequired    = &httpserver.Error{Status: http.StatusBadRequest, Code: "full_name_is_required", Desc: "Full name is required"}
	webErrPhoneNumberIsRequired = &httpserver.Error{Status: http.StatusBadRequest, Code: "phone_number_is_required", Desc: "Phone number is required"}
	webErrInvalidPhoneNumber    = &httpserver.Error{Status: http.StatusBadRequest, Code: "invalid_number", Desc: "Invalid phone number"}
	webErrPasswordIsRequired    = &httpserver.Error{Status: http.StatusBadRequest, Code: "password_is_required", Desc: "Password is required"}
	webErrInvalidPassword       = &httpserver.Error{Status: http.StatusBadRequest, Code: "invalid_password", Desc: "Invalid password. Rules: The password should between 8 and 12 characters, contains at least one uppercase letter, one lowercase letter, one number and one special character"}
	webErrPhoneNumberExists     = &httpserver.Error{Status: http.StatusBadRequest, Code: "phone_number_already_exists", Desc: "Phone number already exists"}
)

func convertToClientErr(err error) error {
	if err == nil {
		return nil
	}

	switch err.Error() {
	case user.ErrPhoneNumberExists.Error():
		return webErrPhoneNumberExists
	default:
		return &httpserver.Error{Status: http.StatusInternalServerError, Code: "internal_error", Desc: err.Error()}
	}
}
