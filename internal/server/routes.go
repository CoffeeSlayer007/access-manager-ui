package server

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed ui
var embeddedFiles embed.FS

func getHandlers() *gin.Engine {
	e := gin.Default()
	configureRoutes(e)
	return e
}

func configureRoutes(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	e.GET("/", func(c *gin.Context) {
		if c.FullPath() == "/" {
			c.Redirect(http.StatusMovedPermanently, "/ui")
		}
	})
	e.StaticFS("/ui", getFileSystem("ui"))
	e.StaticFS("/static", getFileSystem("ui/static"))
}

func getFileSystem(path string) http.FileSystem {

	// Get the build subdirectory as the
	// root directory so that it can be passed
	// to the http.FileServer
	fsys, err := fs.Sub(embeddedFiles, path)
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
