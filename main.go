package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:id`
	Title  string  `json:title`
	Artist string  `json:artist`
	Price  float64 `json:price`
}

var albums = []album{
	{
		ID:     "1",
		Title:  "Blue Trai",
		Artist: "John Coltrane",
		Price:  56.99,
	},
	{
		ID:     "2",
		Title:  "Jeru",
		Artist: "Gerry Mulligan",
		Price:  17.99,
	},
	{
		ID:     "3",
		Title:  "Sarah Vaughan and Clifford Brown",
		Artist: "Sarah Vaughan",
		Price:  39.99,
	},
}

// getAlbums responses with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for index, a := range albums {
		// リソースがあったら204
		if a.ID == id {
			albums = append(albums[:index], albums[index+1:]...)
			c.IndentedJSON(http.StatusNoContent, albums)
			return
		}
	}
	// リソースがなければ404
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// putAlbums update an album from JSON received in the request body
func putAlbumByID(c *gin.Context) {
	var updateAlbum album

	if err := c.BindJSON(&updateAlbum); err != nil {
		return
	}

	id := c.Param("id")

	for index, a := range albums {
		if a.ID == id {
			albums[index] = updateAlbum
			c.IndentedJSON(http.StatusNoContent, albums)
			return
		}
	}

	albums = append(albums, updateAlbum)
	c.IndentedJSON(http.StatusCreated, updateAlbum)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.PUT("/albums/:id", putAlbumByID)
	return router
}

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
}
