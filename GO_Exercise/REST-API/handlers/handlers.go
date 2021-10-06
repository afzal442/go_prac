package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type nwarticle struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
	Review string  `json:"review"`
	Rating float64 `json:"rating"`
}

// article represents data about a record article.
type article struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
	Review string  `json:"review"`
	Rating float64 `json:"rating"`
}

// articles slice to seed record article data.
var articles = []article{
	{ID: "1", Title: "Let's GO", Author: "John Doe", Price: 56.99, Review: "Very Good", Rating: 4},
	{ID: "2", Title: "the three mistakes of my life", Author: "Chetan", Price: 17.99, Review: "Good", Rating: 3.5},
	{ID: "3", Title: "3 idiots", Author: "Bhagat", Price: 39.99, Review: "Excellent", Rating: 5},
}

// getarticles responds with the list of all articles as JSON.
func GetArticles(c *gin.Context) {
	c.JSON(http.StatusOK, articles)
}

func PutData(c *gin.Context) {
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
			// IndentedJSON serializes the article into JSON
			c.IndentedJSON(http.StatusCreated, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "article not found"})

}

// postarticles adds an article from JSON received in the request body.
func PostArticles(c *gin.Context) {
	var newarticle article

	// Call BindJSON to bind the received JSON to
	// newarticle.
	if err := c.BindJSON(&newarticle); err != nil {
		return
	}

	articles = append(articles, newarticle)
	c.IndentedJSON(http.StatusCreated, newarticle)
}

func CreateArticles(w http.ResponseWriter, r *http.Request) {
	var newarticle article
	// Call BindJSON to bind the received JSON to
	// newarticle.
	// if err := json.NewDecoder(r.Body).Decode(&newarticle); err != nil {
	// 	return
	// }

	// get the body of POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newarticle)

	articles = append(articles, newarticle)
	json.NewEncoder(w).Encode(newarticle)
}

// getarticleByID locates the article whose ID value matches the id
// parameter sent by the client, then returns that article as a response.
func GetArtByID(c *gin.Context) {
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

func DeleteData(c *gin.Context) {
	id := c.Param("id")

	for index, a := range articles {
		if a.ID == id {
			articles = append(articles[:index], articles[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "article deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})
}
