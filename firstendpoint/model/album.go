package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var albums = []album{
	{ID: "1", Title: "Familia del artista 01", Artist: "Wlopezob", Year: 2024},
	{ID: "2", Title: "Familia del artista 02", Artist: "Wlopezob", Year: 2025},
	{ID: "3", Title: "Familia del artista 03", Artist: "Wlopezob", Year: 2026},
	{ID: "4", Title: "Familia del artista 04", Artist: "Wlopezob", Year: 2027},
	{ID: "5", Title: "Familia del artista 05", Artist: "Wlopezob", Year: 2037},
}

func GetAlbumns(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum album
	c.BindJSON(&newAlbum)
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	return
}