package api

import (
	"OLXbaSurfer/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// OLXClient communicates with OLX.ba via its API to retrieve requested data.
type OLXClient struct {
	baseURL string
}

// Exported functions

// NewOLXClient constructs and returns an OLXClient object that is ready for use.
func NewOLXClient() Client {
	olxClient := OLXClient{
		baseURL: "https://api.pik.ba/",
	}

	return &olxClient
}

// Exported methods

// SearchArticlesWithQuery runs a query to search for articles with given query and returns the result.
func (olxClient *OLXClient) SearchArticlesWithQuery(query string) ([]models.Article, error) {
	articlesExist := true
	articles := []models.Article{}
	page := 1

	for articlesExist {
		urlString := olxClient.baseURL + "artikli?stranica=" + strconv.Itoa(page) + "&trazilica=" + url.QueryEscape(query)
		resp, err := http.Get(urlString)
		if err != nil {
			return articles, err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return articles, err
		}

		searchResults := models.SearchResult{}
		json.Unmarshal(body, &searchResults)
		if searchResults.Success == false {
			return articles, errors.New("API call not successful")
		}

		if len(searchResults.Articles) == 0 {
			articlesExist = false
		} else {
			articles = append(articles, searchResults.Articles...)
		}

		page++
	}

	return articles, nil
}
