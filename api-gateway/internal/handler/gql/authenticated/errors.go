package authenticated

import (
	"net/http"

	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient/services/user"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/httpserver"
)

var (
	webErrIamIDIsRequired                 = &httpserver.Error{Status: http.StatusBadRequest, Code: "iamid_is_required", Desc: "IamID is required"}
	webErrUserNotFound                    = &httpserver.Error{Status: http.StatusBadRequest, Code: "user_not_found", Desc: "User not found"}
	webErrInvalidUpdateProfileRequestData = &httpserver.Error{Status: http.StatusBadRequest, Code: "invalid_update_profile_request_data", Desc: "Invalid update profile request data"}
)

func convertToClientErr(err error) error {
	if err == nil {
		return nil
	}

	switch err.Error() {
	case user.ErrUserNotFound.Error():
		return webErrUserNotFound
	default:
		return &httpserver.Error{Status: http.StatusInternalServerError, Code: "internal_error", Desc: err.Error()}
	}
}
