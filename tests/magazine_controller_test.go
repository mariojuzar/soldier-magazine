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

func TestPutMagazine(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	magazineSvc := mock.NewMockMagazineService(ctrl)
	magazineCtrl := controller.MagazineController{MagazineService: magazineSvc}

	t.Run("Put Magazine Bind Failed", func(t *testing.T) {
		req := httptest.NewRequest("POST", path.BaseUrl + path.Magazine, bytes.NewBufferString("{}"))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Request = req

		magazineCtrl.PutMagazine(ctx)

		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})

	t.Run("Put Magazine Success", func(t *testing.T) {
		magazineSvc.EXPECT().PutMagazine(gomock.Any()).Return(model.Soldier{}, nil)

		req := httptest.NewRequest("POST", path.BaseUrl + path.Magazine, bytes.NewBufferString("{\"soldier_id\":1,\"magazines\":[{\"max_capacity\":6},{\"max_capacity\":4}]}"))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Request = req

		magazineCtrl.PutMagazine(ctx)

		assert.Equal(t, rec.Code, http.StatusOK)
	})

	t.Run("Put Magazine Failed", func(t *testing.T) {
		magazineSvc.EXPECT().PutMagazine(gomock.Any()).Return(model.Soldier{}, errors.New("error here"))

		req := httptest.NewRequest("POST", path.BaseUrl + path.Magazine, bytes.NewBufferString("{\"soldier_id\":1,\"magazines\":[{\"max_capacity\":6},{\"max_capacity\":4}]}"))
		req.Header.Set("Content-Type", "application/json")

		rec := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rec)
		ctx.Request = req

		magazineCtrl.PutMagazine(ctx)

		assert.Equal(t, rec.Code, http.StatusBadRequest)
	})
}