package http

//request struct
type ShortenURLRequest struct {
	URL string `json:"url" binding:"required"`
}

//error response struct
type ShortenURLErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

//success response struct
type ShortenURLSuccessResponse struct {
	ShortURL string `json:"short_url"`
	Message  string `json:"message"`
}
