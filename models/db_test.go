package models

import (
	"os"
	"testing"
)

func TestGetSetQuery(t *testing.T) {
	database := setupDB("getset", t)

	prevValue := database.GetSetQuery("Test Query #1")
	if prevValue != nil {
		t.Error("Previous value already exists in db. This should not happen.")
	}

	prevValue = database.GetSetQuery("Test Query #1")
	if prevValue == nil {
		t.Error("Previous value should exist in database")
	}

	prevValue = database.GetSetQuery("Test Query #2")
	if prevValue != nil {
		t.Error("Previous value should not exist in db. This should not happen.")
	}
}

func TestStoreArticleAndDoesArticleExist(t *testing.T) {
	database := setupDB("store", t)

	article := Article{
		ID:    1,
		Name:  "Test article #1",
		Price: "100 BAM",
	}

	if database.DoesArticleExist(1) == true {
		t.Error("Article (1) should not exist.")
	}

	database.StoreArticle(&article)

	if database.DoesArticleExist(1) == false {
		t.Error("Article (1) should exist.")
	}
}

func setupDB(dbName string, t *testing.T) Datastore {
	err := os.MkdirAll("/tmp/OLXbaSurfer", os.ModePerm)
	if err != nil {
		t.Error(err)
	}

	var database Datastore
	database, err = NewDB("/tmp/OLXbaSurfer/" + dbName)
	if err != nil {
		t.Error(err)
	}

	database.DropAll()

	return database
}
