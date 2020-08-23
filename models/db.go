package models

import "github.com/dgraph-io/badger"

// Datastore represents interface that will be used to communicate with underlying database.
type Datastore interface {
	DoesArticleExist(id uint64) bool
	StoreArticle(article *Article) error
	GetSetQuery(query string) *string
	DropAll() error
}

// DB is a database object. You should create it with NewDB() exported function.
type DB struct {
	*badger.DB
}

// NewDB creates and returns a pointer to the Badger database.
func NewDB(path string) (*DB, error) {
	options := badger.DefaultOptions(path)
	options.Logger = nil

	db, err := badger.Open(options)
	if err != nil {
		return nil, err
	}

	DB := &DB{
		db,
	}

	return DB, err
}

// GetSetQuery return value if it exists in the database and also sets it.
func (db *DB) GetSetQuery(query string) *string {
	err := db.View(func(txn *badger.Txn) error {
		storedQuery, err := txn.Get([]byte(query))
		if err != nil {
			return err
		}

		err = storedQuery.Value(func(val []byte) error {
			return nil
		})
		return err
	})

	// If it is not already stored, we store it
	if err != nil {
		err = db.Update(func(txn *badger.Txn) error {
			err := txn.Set([]byte([]byte(query)), []byte("1"))
			return err
		})

		return nil
	}

	return &query
}
