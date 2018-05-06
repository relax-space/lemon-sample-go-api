package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"sample-go-api/models"
	"testing"

	"github.com/relax-space/go-kit/test"

	"github.com/labstack/echo"
)

func Test_fruit_GetAll(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	test.Ok(t, handleWithFilter(FruitApiController{}.GetAll, echoApp.NewContext(req, rec)))
	fmt.Println(string(rec.Body.Bytes()))
	fmt.Printf("http status:%v", rec.Result().StatusCode)
}

func Test_fruit_Create(t *testing.T) {
	fruit := &models.Fruit{
		Code: "123",
	}
	b, _ := json.Marshal(fruit)
	req := httptest.NewRequest(echo.POST, "/", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	test.Ok(t, handleWithFilter(FruitApiController{}.Create, echoApp.NewContext(req, rec)))
	fmt.Println(string(rec.Body.Bytes()))
	fmt.Printf("http status:%v", rec.Result().StatusCode)
}

func Test_fruit_Update(t *testing.T) {
	fruit := &models.Fruit{
		Code: "2222",
	}
	b, _ := json.Marshal(fruit)
	req := httptest.NewRequest(echo.POST, "/1", bytes.NewReader(b))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	test.Ok(t, handleWithFilter(FruitApiController{}.Update, echoApp.NewContext(req, rec)))
	fmt.Println(string(rec.Body.Bytes()))
	fmt.Printf("http status:%v", rec.Result().StatusCode)
}

func Test_fruit_Delete(t *testing.T) {
	id := "12"
	req := httptest.NewRequest(echo.POST, "/"+id, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	test.Ok(t, handleWithFilter(FruitApiController{}.Delete, echoApp.NewContext(req, rec)))
	fmt.Println(string(rec.Body.Bytes()))
	fmt.Printf("http status:%v", rec.Result().StatusCode)
}
