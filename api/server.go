package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mariojuzar/soldier-magazine/api/configuration"
	"github.com/mariojuzar/soldier-magazine/api/controller/soldier"
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
	var dbSvc = service.NewDatabaseService()
	_ = dbSvc.Initialize()

	var soldierController = soldier.ControllerSoldier{
		SoldierService: service.NewSoldierService(),
	}

	api := engine.Group(path.BaseUrl)
	{
		// soldier path
		api.POST(path.Soldier, soldierController.CreateSoldier)
		api.PUT(path.Soldier, soldierController.UpdateSoldier)
		api.GET(path.Soldier, soldierController.GetAllSoldier)
		api.GET(path.SoldierById, soldierController.GetSoldier)
		api.DELETE(path.SoldierById, soldierController.DeleteSoldier)
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