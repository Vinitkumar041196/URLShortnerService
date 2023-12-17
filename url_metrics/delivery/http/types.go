package http

//error response struct
type GetTopDomainsErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

//success response struct 
type GetTopDomainsSuccessResponse struct {
	Data    map[string]int `json:"data"`
	Message string         `json:"message"`
}
