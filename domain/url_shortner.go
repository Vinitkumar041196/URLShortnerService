package domain

type URLShortnerService interface {
	ShortenURL(url string) (string, error)
	GetOriginalURL(shortUrl string) (string, error)
}

type URLShortnerRepository interface {
	StoreShortURL(fullURL, shortURL string) error
	GetFullURL(shortURL string) (string, error)
}

const ShortURLPrefix = "/url"