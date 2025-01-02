package backend

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"main/backend/controllers"
)

type Routes struct {
}

func (Routes) Setup(fs embed.FS) *gin.Engine {
	godotenv.Load(".env")
	dev := os.Getenv("APP_ENV") == "development"

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	if dev {
		router.Use(cors.Default())
	}

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		return fmt.Sprintf("%v [GIN] |%s %3d %s| %13v | %15s |%s %-7s %s %#v\n%s",
			param.TimeStamp.Format("2006/01/02 15:04:05"),
			statusColor, param.StatusCode, resetColor,
			param.Latency,
			param.ClientIP,
			methodColor, param.Method, resetColor,
			param.Path,
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	foldersController := controllers.FoldersController{}

	api := router.Group("/api")

	folders := api.Group("/folders")
	folders.GET("", foldersController.GetFolders)

	router.Use(serve("/", embedFolder(fs, "build")))
	router.NoRoute(func(context *gin.Context) { // fallback
		data, err := fs.ReadFile("build/index.html")
		if err != nil {
			context.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		context.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	return router
}

type serveFileSystem interface {
	http.FileSystem
	exists(prefix string, path string) bool
}

func serve(urlPrefix string, fs serveFileSystem) gin.HandlerFunc {
	fileserver := http.FileServer(fs)
	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}
	return func(c *gin.Context) {
		if fs.exists(urlPrefix, c.Request.URL.Path) {
			fileserver.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func embedFolder(fsEmbed embed.FS, targetPath string) serveFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
