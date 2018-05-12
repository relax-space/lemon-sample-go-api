package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"github.com/relax-space/go-kit/base"
	"github.com/relax-space/go-kit/sign"
)

type SignApiController struct {
}

func (d SignApiController) Init(g *echo.Group) {
	g.GET("", d.SignGet)
	g.POST("", d.SignPost)
}

func (SignApiController) SignPost(c echo.Context) error {
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
	}
	defer c.Request().Body.Close()
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(b))
	m := make(map[string]interface{}, 0)
	err = json.Unmarshal(b, &m)
	if err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
	}
	signKey := "hello-sign" // In the operating environment, please keep signKey in a safe place
	var signParam string
	if len(m) != 0 {
		if signObj, ok := m["sign"]; ok {
			signParam = signObj.(string)
			delete(m, "sign")
		}
	}
	//1. check sign
	isOk := sign.CheckMd5Sign(base.JoinMapObject(m), signKey, signParam)
	if isOk != true {
		return ReturnApiFail(c, http.StatusOK, ApiErrorSign, nil, map[string]interface{}{"sign-string": signParam, "sign-param": base.JoinMapObject(m)})
	}
	//2. parse param from body
	var v Book
	if err := c.Bind(&v); err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
	}

	return ReturnApiSucc(c, http.StatusOK, v)
}

func (SignApiController) SignGet(c echo.Context) error {
	signKey := "hello-sign" // In the operating environment, please keep signKey in a safe place
	//1. check sign
	isOk := sign.CheckMd5Sign(base.RemoveFromString(c.QueryString(), "sign"), signKey, c.QueryParam("sign"))
	if isOk != true {
		authors, err := url.PathUnescape(c.QueryParam("authors"))
		if err != nil {
			return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
		}
		return ReturnApiFail(c, http.StatusOK, ApiErrorSign, nil,
			map[string]interface{}{"sign-string": c.QueryParam("sign"), "authors": authors})
	}
	//2. parse param from body
	var v struct {
		Name string `json:"name" query:"name"`
		Sign string `json:"sign" query:"sign"`
	}
	if err := c.Bind(&v); err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
	}
	book := Book{
		Name: v.Name,
		Sign: v.Sign,
	}
	authors, err := url.PathUnescape(c.QueryParam("authors"))
	if err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
	}
	if err = json.Unmarshal([]byte(authors), &book.Authors); err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
	}

	return ReturnApiSucc(c, http.StatusOK, book)
}
