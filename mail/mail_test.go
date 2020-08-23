package mail

import (
	"OLXbaSurfer/models"
	"testing"
)

func TestCreateBodyForArticles(t *testing.T) {
	articles := []models.Article{
		{
			ID:    1,
			Name:  "Test 1",
			Price: "10 BAM",
		},
		{
			ID:    2,
			Name:  "Test 2",
			Price: "15 BAM",
		},
	}

	resultBody := createBodyForArticles(articles)
	desiredBody := `We have found 2 article(s) since last search:
<ul><li><a href="https://www.olx.ba/artikal/1">Test 1</a> with price: 10 BAM</li><li><a href="https://www.olx.ba/artikal/2">Test 2</a> with price: 15 BAM</li></ul>`

	if resultBody != desiredBody {
		t.Errorf("Result mail body is not equal to desired body.\nDesired: %s\nResult: %s", desiredBody, resultBody)
	}
}
