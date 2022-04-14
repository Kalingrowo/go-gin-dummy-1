package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Bintang Di Surga", Artist: "Peterpan", Price: 10.5},
	{ID: "2", Title: "Aay", Artist: "( clone )", Price: 9.99},
	{ID: "3", Title: "kehilangan", Artist: "firman", Price: 7.4},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")
}

// get list albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// store new album
func postAlbum(c *gin.Context) {
	var newAlbum album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

// get specific album
func getAlbumByID(c *gin.Context) {
	albumId := c.Param("id")

	for _, album := range albums {
		if album.ID == albumId {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
