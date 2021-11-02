package main

import (
	modules "go-api-jazz/v1/pkg/modules/albums"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", modules.GetAlbums)
	router.POST("/albums", modules.PostAlbums)
	router.GET("/albums/:id", modules.GetAlbumById)

	router.Run("localhost:8080")
}
