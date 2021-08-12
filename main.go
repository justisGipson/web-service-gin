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

// getAlbums responds with list of all albums as json
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
