package main

import (
	"log"
	"github.com/jmoiron/sqlx"
	"github.com/gin-gonic/gin"
)



// getAlbums responds witht he list of all albums as JSON
func getAlbums(c *gin.Context, db sqlx.DB) {
	rows, err := db.Query("SELECT * FROM albums")
	if er != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// c.IndentedJSON(http.StatusOK, albums)
}
func main {
	db, err := sqlx.Open("postgres", )
}