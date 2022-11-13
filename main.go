package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	globals "github.com/satouyuuki/golang-simple-api/globals"
	helpers "github.com/satouyuuki/golang-simple-api/helpers"
	middleware "github.com/satouyuuki/golang-simple-api/middleware"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type account struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
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

func login(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Please logout first"})
		return
	}

	var requestBody account
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	username := requestBody.Username
	password := requestBody.Password

	log.Println(username, password)

	if helpers.EmptyUserPass(username, password) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Parameters can't be empty"})
		return
	}

	if !helpers.CheckUserPass(username, password) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Incorrect username or password"})
		return
	}

	session.Set(globals.Userkey, username)
	if err := session.Save(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to save session"})
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user == nil {
		log.Println("Invalid session token")
		return
	}
	session.Delete(globals.Userkey)
	if err := session.Save(); err != nil {
		log.Println("Failed to save session:", err)
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func index(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "This is an index page...",
		"user":    user,
	})
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Setup the cookie store for session management
	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	// Private group, require authentication to access
	public := router.Group("/")
	public.POST("/login", login)
	public.GET("/", index)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	private.GET("/albums", getAlbums)
	private.GET("/albums/:id", getAlbumByID)
	private.POST("/albums", postAlbums)
	private.DELETE("/albums/:id", deleteAlbumByID)
	private.PUT("/albums/:id", putAlbumByID)
	private.POST("/logout", logout)
	return router
}

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
}
