package controllers

import (
	"net/http"
	"sample-go-api/models"
	"strconv"

	"github.com/labstack/echo"
)

type FruitApiController struct {
}

func (d FruitApiController) Init(g *echo.Group) {
	g.GET("", d.GetAll)
	g.GET("/:id", d.GetOne)
	g.PUT("/:id", d.Update)
	g.POST("", d.Create)
	g.DELETE("/:id", d.Delete)
}

func (FruitApiController) GetAll(c echo.Context) error {
	var v SearchInput
	if err := c.Bind(&v); err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
	}
	if v.MaxResultCount == 0 {
		v.MaxResultCount = DefaultMaxResultCount
	}

	totalCount, items, err := models.Fruit{}.GetAll(c.Request().Context(), v.Sortby, v.Order, v.SkipCount, v.MaxResultCount)
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, ApiErrorDB, err)
	}
	if len(items) == 0 {
		return ReturnApiFail(c, http.StatusNotFound, ApiErrorNotFound, err)
	}
	return ReturnApiListSucc(c, http.StatusOK, totalCount, items)
}

func (d FruitApiController) Create(c echo.Context) error {
	var v models.Fruit
	if err := c.Bind(&v); err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
	}
	if err := c.Validate(v); err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
	}
	return d.create(c, &v)
}

func (FruitApiController) create(c echo.Context, fruit *models.Fruit) error {
	has, _, err := models.Fruit{}.GetById(c.Request().Context(), fruit.Id)
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, ApiErrorDB, err)
	}
	if has {
		return ReturnApiFail(c, http.StatusOK, ApiErrorHasExist, err)
	}
	affectedRow, err := fruit.Create(c.Request().Context())
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, ApiErrorDB, err)
	}
	if affectedRow == 0 {
		return ReturnApiFail(c, http.StatusOK, ApiErrorNotChanged, err)
	}
	return ReturnApiSucc(c, http.StatusCreated, fruit)
}

func (FruitApiController) GetOne(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err, map[string]interface{}{"id": c.Param("id")})
	}
	has, v, err := models.Fruit{}.GetById(c.Request().Context(), id)
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, ApiErrorDB, err)
	}
	if !has {
		return ReturnApiFail(c, http.StatusNotFound, ApiErrorNotFound, nil)
	}
	return ReturnApiSucc(c, http.StatusOK, v)
}

func (d FruitApiController) Update(c echo.Context) error {
	var v models.Fruit
	if err := c.Bind(&v); err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
	}
	if err := c.Validate(&v); err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err, map[string]interface{}{"id": c.Param("id")})
	}
	v.Id = id
	return d.update(c, &v)
}

func (FruitApiController) update(c echo.Context, v *models.Fruit) error {
	has, _, err := models.Fruit{}.GetById(c.Request().Context(), v.Id)
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, ApiErrorDB, err)
	}
	if !has {
		return ReturnApiFail(c, http.StatusNotFound, ApiErrorNotFound, err)
	}
	affectedRow, err := v.Update(c.Request().Context(), v.Id)
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, ApiErrorDB, err)
	}
	if affectedRow == 0 {
		return ReturnApiFail(c, http.StatusOK, ApiErrorNotChanged, err)
	}
	return ReturnApiSucc(c, http.StatusNoContent, nil)
}

func (d FruitApiController) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err, map[string]interface{}{"id": c.Param("id")})
	}
	return d.delete(c, id)
}

func (FruitApiController) delete(c echo.Context, id int64) error {
	has, _, err := models.Fruit{}.GetById(c.Request().Context(), id)
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, ApiErrorDB, err)
	}
	if !has {
		return ReturnApiFail(c, http.StatusNotFound, ApiErrorNotFound, err)
	}
	affectedRow, err := models.Fruit{}.Delete(c.Request().Context(), id)
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, ApiErrorDB, err)
	}
	if affectedRow == 0 {
		return ReturnApiFail(c, http.StatusOK, ApiErrorNotChanged, err)
	}
	return ReturnApiSucc(c, http.StatusNoContent, nil)
}
