package welcomertr

import "github.com/gin-gonic/gin"

func welcomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}
