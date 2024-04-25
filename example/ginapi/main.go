package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/SupTarr/intensive_go_basic_workshop/_exercises/thai_id"
)

type VerifyIdRequest struct {
	ID string `json:"id" binding:"required"`
}

type VerifyIdResponse struct {
	Valid bool `json:"valid"`
}

type VerifyIdErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/thai/ids/verify", func(c *gin.Context) {
		var json VerifyIdRequest
		err := c.BindJSON(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest, VerifyIdErrorResponse{
				Message: err.Error(),
			})

			return
		}

		err = thai_id.Validate(json.ID)
		if err != nil {
			c.JSON(http.StatusOK, VerifyIdResponse{
				Valid: false,
			})

			return
		}

		c.JSON(http.StatusOK, VerifyIdResponse{
			Valid: true,
		})
	})

	r.Run()
}
