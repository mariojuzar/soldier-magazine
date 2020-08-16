package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mariojuzar/soldier-magazine/api/entity/rest-web/request"
	baseResponse "github.com/mariojuzar/soldier-magazine/api/entity/rest-web/response"
	"github.com/mariojuzar/soldier-magazine/api/libraries/util"
	"github.com/mariojuzar/soldier-magazine/api/service"
	"net/http"
	"time"
)

type SoldierController struct {
	service.SoldierService
}

func (sc *SoldierController) CreateSoldier(c *gin.Context) {
	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	var soldierRequest request.SoldierCreateRequest
	err := c.Bind(&soldierRequest)

	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		soldier, err := sc.SoldierService.CreateSoldier(soldierRequest.Name)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusCreated
			response.Message = http.StatusText(http.StatusCreated)
			response.Data = soldier

			c.JSON(http.StatusCreated, response)
		}
	}
}

func (sc *SoldierController) GetSoldier(c *gin.Context) {
	id := util.StrToUint(c.Params.ByName("id"))

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if id == 0 {
		response.Code = http.StatusBadRequest
		response.Message = "Required id"

		c.JSON(http.StatusBadRequest, response)
	} else {
		soldier, err := sc.SoldierService.GetSoldier(id)

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

func (sc *SoldierController) UpdateSoldier(c *gin.Context) {
	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	var soldierRequest request.SoldierUpdateRequest
	err := c.Bind(&soldierRequest)

	if err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()

		c.JSON(http.StatusBadRequest, response)
	} else {
		soldier, err := sc.SoldierService.UpdateSoldier(soldierRequest.ID, soldierRequest.Name)

		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()

			c.JSON(http.StatusBadRequest, response)
		} else {
			response.Code = http.StatusCreated
			response.Message = http.StatusText(http.StatusCreated)
			response.Data = soldier

			c.JSON(http.StatusCreated, response)
		}
	}
}

func (sc *SoldierController) GetAllSoldier(c *gin.Context) {
	soldiers, err := sc.SoldierService.GetAllSoldier()

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if err != nil {
		response.Code = http.StatusInternalServerError
		response.Message = err.Error()

		c.JSON(http.StatusInternalServerError, response)
	} else {
		response.Code = http.StatusOK
		response.Message = http.StatusText(http.StatusOK)
		response.Data = soldiers

		c.JSON(http.StatusOK, response)
	}
}

func (sc *SoldierController) DeleteSoldier(c *gin.Context) {
	id := util.StrToUint(c.Params.ByName("id"))

	var response = &baseResponse.BaseResponse{
		ServerTime:	time.Now(),
	}

	if id == 0 {
		response.Code = http.StatusBadRequest
		response.Message = "Required id"

		c.JSON(http.StatusBadRequest, response)
	} else {
		soldier, err := sc.SoldierService.DeleteSoldier(id)

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