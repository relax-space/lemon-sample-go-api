package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/pangpanglabs/goutils/behaviorlog"

	"github.com/fatih/structs"
	"github.com/labstack/echo"
)

const (
	FlashName      = "flash"
	FlashSeparator = ";"
)

type ApiResult struct {
	Result  interface{} `json:"result"`
	Success bool        `json:"success"`
	Error   ApiError    `json:"error"`
}

type ApiError struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

type ArrayResult struct {
	Items      interface{} `json:"items"`
	TotalCount int64       `json:"totalCount"`
}

var (
	// System Error
	ApiErrorSystem             = ApiError{Code: 10001, Message: "System Error"}
	ApiErrorServiceUnavailable = ApiError{Code: 10002, Message: "Service unavailable"}
	ApiErrorRemoteService      = ApiError{Code: 10003, Message: "Remote service error"}
	ApiErrorIPLimit            = ApiError{Code: 10004, Message: "IP limit"}
	ApiErrorPermissionDenied   = ApiError{Code: 10005, Message: "Permission denied"}
	ApiErrorIllegalRequest     = ApiError{Code: 10006, Message: "Illegal request"}
	ApiErrorHTTPMethod         = ApiError{Code: 10007, Message: "HTTP method is not suported for this request"}
	ApiErrorParameter          = ApiError{Code: 10008, Message: "Parameter error"}
	ApiErrorMissParameter      = ApiError{Code: 10009, Message: "Miss required parameter"}
	ApiErrorDB                 = ApiError{Code: 10010, Message: "DB error, please contact the administator"}
	ApiErrorTokenInvaild       = ApiError{Code: 10011, Message: "Token invaild"}
	ApiErrorMissToken          = ApiError{Code: 10012, Message: "Miss token"}
	ApiErrorVersion            = ApiError{Code: 10013, Message: "API version %s invalid"}
	ApiErrorNotFound           = ApiError{Code: 10014, Message: "Resource has not found"}
	ApiErrorHasExist           = ApiError{Code: 10015, Message: "Resource has existed"}
	ApiErrorNotChanged         = ApiError{Code: 10016, Message: "Resource has not changed"}
	ApiErrorSign               = ApiError{Code: 10017, Message: "Sign invaild"}

	// Business Error
	ApiErrorUserNotExists = ApiError{Code: 20001, Message: "User does not exists"}
	ApiErrorPassword      = ApiError{Code: 20002, Message: "Password error"}
)

func ReturnApiFail(c echo.Context, status int, apiError ApiError, err error, v ...map[string]interface{}) error {
	logContext := behaviorlog.FromCtx(c.Request().Context())
	if logContext != nil {
		if err != nil {
			logContext.WithError(err)
		}
		if len(v) > 0 {
			logContext.WithBizAttrs(v[0])
		}
	}

	str := ""
	if err != nil {
		str = err.Error()
	}
	return c.JSON(status, ApiResult{
		Success: false,
		Error: ApiError{
			Code:    apiError.Code,
			Message: apiError.Message,
			Details: str,
		},
	})
}

func ReturnApiSucc(c echo.Context, status int, result interface{}) error {
	if status == 204 {
		return c.NoContent(status)
	}

	return c.JSON(status, ApiResult{
		Success: true,
		Result:  result,
	})
}
func ReturnApiListSucc(c echo.Context, status int, totalCount int64, items interface{}) error {
	if status == 204 {
		return c.NoContent(status)
	}
	return c.JSON(status, ApiResult{
		Success: true,
		Result:  ArrayResult{TotalCount: totalCount, Items: items},
	})
}

func setFlashMessage(c echo.Context, m map[string]string) {
	var flashValue string
	for key, value := range m {
		flashValue += "\x00" + key + "\x23" + FlashSeparator + "\x23" + value + "\x00"
	}

	c.SetCookie(&http.Cookie{
		Name:  FlashName,
		Value: url.QueryEscape(flashValue),
	})
}
func getFlashMessage(c echo.Context) map[string]string {
	cookie, err := c.Cookie(FlashName)
	if err != nil {
		return nil
	}

	m := map[string]string{}

	v, _ := url.QueryUnescape(cookie.Value)
	vals := strings.Split(v, "\x00")
	for _, v := range vals {
		if len(v) > 0 {
			kv := strings.Split(v, "\x23"+FlashSeparator+"\x23")
			if len(kv) == 2 {
				m[kv[0]] = kv[1]
			}
		}
	}
	//read one time then delete it
	c.SetCookie(&http.Cookie{
		Name:   FlashName,
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})

	return m
}

type UrlInfo struct {
	ControllerName string
	ApiName        string
	Method         string //GET,POST
	Uri            string
	ResponseStatus interface{}
	Struct         interface{}
	Err            error
}

func PrintApiBehaviorError(c context.Context, urlInfo UrlInfo) {
	logContext := behaviorlog.FromCtx(c)
	if logContext != nil {
		logClone := logContext.Clone()
		if urlInfo.Err != nil {
			logClone.WithError(urlInfo.Err)
		}
		logClone.Controller = urlInfo.ControllerName
		logClone.Params = map[string]interface{}{}
		param := make(map[string]interface{}, 0)
		if urlInfo.Struct != nil && !reflect.ValueOf(urlInfo.Struct).IsNil() {
			s := structs.New(urlInfo.Struct)
			s.TagName = "json"
			param = s.Map()
		}
		var statusCode int
		switch t := urlInfo.ResponseStatus.(type) {
		case int:
			statusCode = t
		case *http.Response:
			statusCode = t.StatusCode
		}

		logClone.WithCallURLInfo(
			urlInfo.Method,
			urlInfo.Uri,
			param,
			statusCode,
		).Log(urlInfo.ApiName)
		logContext.Params = map[string]interface{}{}
	}
}
func routeParse(c echo.Context) (method, version string, err error) {
	b, err := ioutil.ReadAll(c.Request().Body)
	c.Request().Body.Close()
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(b))
	var reqDto struct {
		Method  string `json:"method"`
		Version string `json:"version"`
	}
	err = json.Unmarshal(b, &reqDto)
	if err != nil {
		return
	}
	method = reqDto.Method
	version = reqDto.Version
	if len(version) == 0 {
		version = "V2"
	}
	return
}
