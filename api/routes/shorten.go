package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raj-gupta/url-shortner-redis-docker/helpers"
)

type request struct {
	URl         string        `json:"url"`
	Expiry      time.Duration `json:"time"`
	CustomShort string        `json:"custom_short"`
}

type response struct {
	URl             string `json:"url"`
	CustomShort     string `json:"custom_short"`
	Expiry          string `json:"expiry"`
	XRateRemaining  string `json:"rate_remaining"`
	XRateLimitReset string `json:"rate_limit"`
}

func ShortenUrl(ctx *gin.Context) {
	body := new(request)
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//implement the rate limiting

	//check if the input is an actual url

	// if !govalidator.IsUrl(body.URl) {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"error": "Invalid Url",
	// 	})
	// 	return
	// }

	//check for domain error
	if !helpers.RemoveDomainError(body.URl) {
		ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
			"error": "Invalid Url",
		})
		return
	}

	//enforce http, SSL

	body.URl = helpers.EnforceHTTP(body.URl)

}
