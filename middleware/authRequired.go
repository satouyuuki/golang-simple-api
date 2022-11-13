package middleware

import (
	"github.com/gin-contrib/sessions"
	globals "github.com/satouyuuki/golang-simple-api/globals"

	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(globals.Userkey)
	if user == nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "accessToken invalid or expired"})
		c.Abort()
	} else {
		c.Next()
	}
}
