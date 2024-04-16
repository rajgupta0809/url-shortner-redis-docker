package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/raj-gupta/url-shortner-redis-docker/database"
)

func ResolveUrl(ctx *gin.Context) {
	params := ctx.Params
	url, _ := params.Get("url")

	r := database.CreateClient(0)
	defer r.Close()

	value := r.Get(url)
	if value.Err() == redis.Nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Short Url Not Found",
		})
		return
	}

	ctx.Redirect(301, value.String())
}
