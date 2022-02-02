package news

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllArticlesHandler(c *gin.Context) {
	c.JSON(http.StatusOK, AllArticles())
}

func ArticlesFromHandler(c *gin.Context) {
	if articles, err := ArticlesFrom(c.Param("src")); err != nil {
		c.JSON(http.StatusOK, articles)
	} else {
		c.AbortWithError(http.StatusNotFound, err)
	}
}
