package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	"github.com/SupTarr/intensive_go_basic_workshop/exercises/thai_id"
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

type ProgrammingLanguage struct {
	Name     string `json:"name"`
	ImageUrl string `json:"imageUrl"`
}

func main() {
	db, err := sql.Open("sqlite3", "../../languages.sqlite")
	if err != nil {
		log.Fatal(err)
	}

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

	r.GET("/languages", func(c *gin.Context) {
		rows, err := db.Query("SELECT name, imageUrl FROM languages")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
		}
		defer rows.Close()

		var response []ProgrammingLanguage
		for rows.Next() {
			var name, imageUrl string
			if err := rows.Scan(&name, &imageUrl); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"message": err.Error(),
				})
			}

			response = append(response, ProgrammingLanguage{
				Name:     name,
				ImageUrl: imageUrl,
			})
		}

		c.JSON(http.StatusOK, response)
	})

	r.Run()
}
