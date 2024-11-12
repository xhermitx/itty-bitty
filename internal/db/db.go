package db

import (
	"github.com/xhermitx/itty-bitty/internal/utils"
)

type Pair struct {
	OriginalURL string
	ShortURL    string
}

type DB struct {
	pairs []Pair
}

func New() *DB {
	return &DB{
		pairs: []Pair{},
	}
}

func (db *DB) Save(originalURL, shortURL string) error {

	for _, pair := range db.pairs {
		if shortURL == pair.ShortURL {
			return utils.ErrAlreadyExists
		}
	}

	db.pairs = append(db.pairs, Pair{
		OriginalURL: originalURL,
		ShortURL:    shortURL,
	})

	return nil
}

//func (db *DB) Delete(url string) error {
//	if len(db.pairs) == 0 {
//		return utils.ErrEmptyDB
//	}
//	// Delete the instance
//	return nil
//}

func (db *DB) GetOriginalURL(shortURL string) (string, error) {

	if len(db.pairs) == 0 {
		return "", utils.ErrEmptyDB
	}

	for _, pair := range db.pairs {
		if pair.ShortURL == shortURL {
			return pair.OriginalURL, nil
		}
	}

	return "", utils.ErrNotFound
}

func (db *DB) GetShortURL(originalURL string) (string, error) {

	if len(db.pairs) == 0 {
		return "", utils.ErrEmptyDB
	}

	for _, pair := range db.pairs {
		if pair.OriginalURL == originalURL {
			return pair.ShortURL, nil
		}
	}

	return "", utils.ErrNotFound
}
