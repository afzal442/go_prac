package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type nwarticle struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// article represents data about a record article.
type article struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// articles slice to seed record article data.
var articles = []article{
	{ID: "1", Title: "Let's GO", Author: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "the three mistakes of my life", Author: "Chetan", Price: 17.99},
	{ID: "3", Title: "3 idiots", Author: "Bhagat", Price: 39.99},
}

// getarticles responds with the list of all articles as JSON.
func getArticles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, articles)
}

func putData(c *gin.Context) {
	var newarticle nwarticle
	id := c.Param("id")

	// Call BindJSON to bind the received JSON to
	// newarticle.
	if err := c.BindJSON(&newarticle); err != nil {
		return
	}

	for _, a := range articles {
		if a.ID == id {
			a.Title = newarticle.Title
			a.Author = newarticle.Author
			a.Price = newarticle.Price
			c.IndentedJSON(http.StatusCreated, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})

}

// postarticles adds an article from JSON received in the request body.
func postArticles(c *gin.Context) {
	var newarticle article

	// Call BindJSON to bind the received JSON to
	// newarticle.
	if err := c.BindJSON(&newarticle); err != nil {
		return
	}

	articles = append(articles, newarticle)
	c.IndentedJSON(http.StatusCreated, newarticle)
}

// getarticleByID locates the article whose ID value matches the id
// parameter sent by the client, then returns that article as a response.
func getArtByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of articles, looking for
	// an article whose ID value matches the parameter.
	for _, a := range articles {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})
}
