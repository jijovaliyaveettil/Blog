package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserUpdate(c *gin.Context) {

	// var req models.User

	user_id := c.Param("id")

	if user_id == "123" {
		fmt.Println("Update the id:", user_id)

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Inside Update",
	})
}
