package http

type ShortenURLRequest struct {
	URL string `json:"url" binding:"required"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type ShortenURLSuccessResponse struct {
	ShortURL string `json:"short_url"`
	Message  string `json:"message"`
}
