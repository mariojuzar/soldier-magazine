package tests

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/mariojuzar/soldier-magazine/api/controller"
	"github.com/mariojuzar/soldier-magazine/api/entity/model"
	"github.com/mariojuzar/soldier-magazine/api/entity/path"
	"github.com/mariojuzar/soldier-magazine/tests/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestCreateSoldier(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	soldierService := mock.NewMockSoldierService(ctrl)
	soldierCtrl := controller.SoldierController{SoldierService: soldierService}

	t.Run("Create Bind Failed", func(t *testing.T) {
		req := httptest.NewRequest("POST", path.BaseUrl + path.Soldier, bytes.NewBufferString("{}"))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Request = req

		soldierCtrl.CreateSoldier(ctx)

		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})

	t.Run("Create Success", func(t *testing.T) {
		soldierService.EXPECT().CreateSoldier("test").Return(model.Soldier{}, nil)

		req := httptest.NewRequest("POST", path.BaseUrl + path.Soldier, bytes.NewBufferString("{\"name\" : \"test\"}"))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Request = req

		soldierCtrl.CreateSoldier(ctx)

		assert.Equal(t, rec.Code, http.StatusCreated)
	})

	t.Run("Create Failed", func(t *testing.T) {
		soldierService.EXPECT().CreateSoldier(gomock.Any()).Return(model.Soldier{}, errors.New("error here"))

		req := httptest.NewRequest("POST", path.BaseUrl + path.Soldier, bytes.NewBufferString("{\"name\" : \"test\"}"))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Request = req

		soldierCtrl.CreateSoldier(ctx)

		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})

}