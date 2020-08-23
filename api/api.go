package api

import "OLXbaSurfer/models"

// Client is a API client that will retrieve articles from OLX.ba
type Client interface {
	SearchArticlesWithQuery(query string) ([]models.Article, error)
}
