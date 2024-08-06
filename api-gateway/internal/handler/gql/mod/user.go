package mod

type SignupRequest struct {
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type SignupResponse struct {
	IamID int64 `json:"iamID"`
}
