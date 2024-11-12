package controllers

import (
	"html/template"
	"net/http"
)

// UrlService defines a contract for a URL service
type UrlService interface {
	ShortenURL(originalURL string) (string, error)
	ValidateURL(originalURL string) (string, error)
	RetrieveOriginalURL(shortURL string) (string, error)
}

type Controller struct {
	svc UrlService
}

func NewController(svc UrlService) *Controller {
	return &Controller{
		svc: svc,
	}
}

func (c *Controller) Shortener(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	inputURL := r.FormValue("url")
	validatedURL, err := c.svc.ValidateURL(inputURL)
	if err != nil {
		http.Error(w, "invalid URL", http.StatusBadRequest)
		return
	}

	// ittyBitty is the shortened URL
	ittyBitty, err := c.svc.ShortenURL(validatedURL)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"ShortURL": ittyBitty,
	}

	t, err := template.ParseFiles("internal/views/shorten.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err = t.Execute(w, data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) Redirect(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:] // Added [1:] slice to remove "/"

	originalURL, err := c.svc.RetrieveOriginalURL(shortURL)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusPermanentRedirect)
}
