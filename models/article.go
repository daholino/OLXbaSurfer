package models

import (
	"OLXbaSurfer/helpers"
	"encoding/json"

	"github.com/dgraph-io/badger"
)

// Article is a listing on OLX.ba.
type Article struct {
	ID    uint64 `json:"id"`
	Name  string `json:"naslov"`
	Price string `json:"cijena"`
}

// StoreArticle stores article to underyling database.
func (db *DB) StoreArticle(article *Article) error {
	articleData, err := json.Marshal(article)
	if err != nil {
		return err
	}

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(helpers.I64tob(article.ID)), articleData)
		return err
	})

	return err
}

// DoesArticleExist checks if the article exist in database.
func (db *DB) DoesArticleExist(id uint64) bool {
	err := db.View(func(txn *badger.Txn) error {
		article, err := txn.Get(helpers.I64tob(id))
		if err != nil {
			return err
		}

		err = article.Value(func(val []byte) error {
			return nil
		})
		return err
	})

	return err == nil
}
