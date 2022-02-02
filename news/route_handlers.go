package news

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllArticlesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, AllArticles())
}

func ArticlesFromHandler(c *gin.Context) {
	src := c.Param("src")

	_, found := readers[src]

	if !found {
		c.String(http.StatusNotFound, "404 page not found")
		return
	}

	if articles, err := ArticlesFrom("dr"); err == nil {
		c.JSON(http.StatusOK, articles)
	} else {
		c.String(http.StatusInternalServerError, "500 internal server error")
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}
