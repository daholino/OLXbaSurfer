package models

// SearchResult represents search result response from OLX.ba API
type SearchResult struct {
	Success      bool      `json:"success"`
	NumOfResults uint32    `json:"broj_rezultata"`
	Articles     []Article `json:"artikli"`
}
