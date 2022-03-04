package main

import (
	"net/http"

	"github.com/chrisUsick/safertogether-server/search"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	search, err := search.NewSearch()
	if err != nil {
		panic(err)
	}

	//new template engine
	router.HTMLRender = ginview.New(goview.Config{
		Root:         "views",
		Extension:    ".html",
		Master:       "layouts/main",
		Partials:     []string{},
		DisableCache: true,
		Delims:       goview.Delims{Left: "{{", Right: "}}"},
	})

	router.GET("/", func(c *gin.Context) {
		localize := NewLocalizerFromContext(c)
		c.HTML(http.StatusOK, "index", gin.H{
			"title": localize("brand_title"),
			"i":     localize,
		})
	})

	router.POST("/search", func(c *gin.Context) {
		searchText := c.PostForm("search")
		results, err := search.SearchText(searchText)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		localize := NewLocalizerFromContext(c)
		c.HTML(http.StatusOK, "search", gin.H{
			"title":   localize("search_title"),
			"i":       localize,
			"results": results,
		})
	})

	router.GET("/page", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "Page file title!!"})
	})

	router.Run("127.0.0.1:8080")
}
