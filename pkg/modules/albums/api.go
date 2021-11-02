package modules

import (
	models "go-api-jazz/v1/pkg/models/album"

	"net/http"

	"github.com/gin-gonic/gin"
)

var albumsList = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Giant Steps", Artist: "John Coltrane", Price: 89.89},
	{ID: "3", Title: "Songs for my Father", Artist: "Horace Silver", Price: 56.99}}

func GetAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albumsList)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albumsList {
		if a.ID == id {

			c.IndentedJSON(http.StatusOK, a)
		}
	}
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albumsList = append(albumsList, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
