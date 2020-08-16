package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mariojuzar/soldier-magazine/api/configuration"
	"github.com/mariojuzar/soldier-magazine/api/entity/path"
	"github.com/mariojuzar/soldier-magazine/api/entity/rest-web/response"
	"github.com/mariojuzar/soldier-magazine/api/service"
	"net/http"
	"time"
)

func Run() *gin.Engine {
	engine := gin.Default()
	engine.RedirectTrailingSlash = false

	configuration.Initialize()
	service.Initialize()

	api := engine.Group(path.BaseUrl)
	{
		api.POST(path.Soldier, nil)
	}

	engine.NoRoute(func(context *gin.Context) {
		var resp = &response.BaseResponse{
			ServerTime:	time.Now(),
		}

		resp.Code = http.StatusNotFound
		resp.Message = "Route not found"

		context.JSON(http.StatusNotFound, resp)
	})

	return engine
}