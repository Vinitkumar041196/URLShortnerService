package domain

type URLShortenerService interface {
	ShortenURL(url string) (string, error)
	GetOriginalURL(shortUrl string) (string, error)
}

type URLShortenerRepository interface {
	StoreShortURL(fullURL, shortURL string) error
	GetFullURL(shortURL string) (string, error)
}

const ShortURLFormat = "%s://%s/v1/%s"
