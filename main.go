package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album data representation
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Kind of Blue", Artist: "Miles Davis", Price: 20.99},
	{ID: "3", Title: "Enter the Wu-Tang", Artist: "Wu-Tang Clan", Price: 17.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums responds with list of all albums as json
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds album from json from req body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// call BindJSON to bind received json to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add newAlbum to slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID fetches album with mathing id sent by client
// returns album as response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// loop over list of albums, searching for id that matches
	// parameter sent
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
