package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mariojuzar/soldier-magazine/api/entity/rest-web/request"
	baseResponse "github.com/mariojuzar/soldier-magazine/api/entity/rest-web/response"
	"github.com/mariojuzar/soldier-magazine/api/service"
	"net/http"
	"time"
)

type MagazineController struct {
	service.MagazineService
}

func (mc *MagazineController) PutMagazine(c *gin.Context) {
	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	var putMagazineRequest request.PutMagazineRequest
	err := c.Bind(&putMagazineRequest)

	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		soldier, err := mc.MagazineService.PutMagazine(putMagazineRequest)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusOK
			response.Message = http.StatusText(http.StatusOK)
			response.Data = soldier

			c.JSON(http.StatusOK, response)
		}
	}
}

func (mc *MagazineController) LoadMagazineRandomly(c *gin.Context) {
	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	var loadMagazineRequest request.LoadMagazineRequest
	err := c.Bind(&loadMagazineRequest)

	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		verified, err := mc.MagazineService.LoadMagazineRandomly(loadMagazineRequest)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusOK
			response.Message = http.StatusText(http.StatusOK)
			response.Data = verified

			c.JSON(http.StatusOK, response)
		}
	}
}