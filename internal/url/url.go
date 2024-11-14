package url

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/xhermitx/itty-bitty/internal/utils"
	"net/url"
	"regexp"
)

type DB interface {
	Save(originalURL, shortURL string) error
	// GetOriginalURL Delete(url string) error
	GetOriginalURL(shortURL string) (string, error)
	GetShortURL(originalURL string) (string, error)
}

type Service struct {
	db DB
}

func NewService(db DB) *Service {
	return &Service{
		db: db,
	}
}

// ShortenURL Implement the Shortening Logic
func (s *Service) ShortenURL(originalURL string) (string, error) {
	h := sha256.New()
	h.Write([]byte(originalURL))
	hash := h.Sum(nil)
	shortURL := hex.EncodeToString(hash)[:8]

	if _, err := s.db.GetShortURL(originalURL); err == nil {
		return "", utils.ErrAlreadyExists
	}

	// Using only the first 8 characters
	if err := s.db.Save(originalURL, shortURL); err != nil {
		return "", err
	}
	return shortURL, nil
}

// ValidateURL validates the URL format
func (s *Service) ValidateURL(originalURL string) (string, error) {
	parsedURL, err := url.Parse(originalURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return "", utils.ErrInvalidURL
	}

	// Define a regular expression to check for a top-level domain
	tldPattern := `\.[a-zA-Z]{2,}$`
	isValidHost := regexp.MustCompile(tldPattern).MatchString(parsedURL.Host)

	// Ensure the URL has a valid top-level domain
	if !isValidHost {
		return "", utils.ErrInvalidURL
	}

	return parsedURL.String(), nil
}

// RetrieveOriginalURL gets the original URL from the DB
func (s *Service) RetrieveOriginalURL(shortURL string) (string, error) {
	shortURl, err := s.db.GetOriginalURL(shortURL)
	if err != nil {
		return "", err
	}
	return shortURl, nil
}
