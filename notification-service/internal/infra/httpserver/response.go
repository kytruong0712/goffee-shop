package httpserver

// Response represents response struct
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
