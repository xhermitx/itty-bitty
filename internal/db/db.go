package db

import (
	"github.com/xhermitx/itty-bitty/internal/utils"
)

type Store struct {
	OriginalURL string
	ShortURL    string
}

type DB struct {
	stores []Store
}

func New() *DB {
	return &DB{
		stores: []Store{},
	}
}

func (db *DB) Save(originalURL, shortURL string) error {
	for _, pair := range db.stores {
		if shortURL == pair.ShortURL {
			return utils.ErrAlreadyExists
		}
	}

	db.stores = append(db.stores, Store{
		OriginalURL: originalURL,
		ShortURL:    shortURL,
	})
	return nil
}

func (db *DB) GetOriginalURL(shortURL string) (string, error) {
	if len(db.stores) == 0 {
		return "", utils.ErrEmptyDB
	}

	for _, pair := range db.stores {
		if pair.ShortURL == shortURL {
			return pair.OriginalURL, nil
		}
	}
	return "", utils.ErrNotFound
}

func (db *DB) GetShortURL(originalURL string) (string, error) {
	if len(db.stores) == 0 {
		return "", utils.ErrEmptyDB
	}

	for _, pair := range db.stores {
		if pair.OriginalURL == originalURL {
			return pair.ShortURL, nil
		}
	}
	return "", utils.ErrNotFound
}
