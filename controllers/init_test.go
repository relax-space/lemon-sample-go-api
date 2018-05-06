package controllers

import (
	"context"
	"os"
	"sample-go-api/factory"

	"github.com/asaskevich/govalidator"
	"github.com/pangpanglabs/goutils/echomiddleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

var (
	echoApp          *echo.Echo
	handleWithFilter func(handlerFunc echo.HandlerFunc, c echo.Context) error
	ctx              context.Context
)

func init() {
	db, err := initDB("mysql", os.Getenv("Fruit_CONN"))
	if err != nil {
		panic(err)
	}

	echoApp = echo.New()
	echoApp.Validator = &Validator{}
	configMap := map[string]interface{}{
		"key": os.Getenv("XXX"),
	}
	setContextValueMiddleware := setContextValue(&configMap, db)
	handleWithFilter = func(handlerFunc echo.HandlerFunc, c echo.Context) error {
		return setContextValueMiddleware(handlerFunc)(c)
	}
}

func setContextValue(configMap *map[string]interface{}, db *xorm.Engine) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			reqContext := context.WithValue(req.Context(), echomiddleware.ContextDBName, db)
			reqContext = context.WithValue(reqContext, factory.ContextConfigName, configMap)
			c.SetRequest(req.WithContext(reqContext))
			return next(c)
		}
	}
}

func initDB(driver, connection string) (*xorm.Engine, error) {
	db, err := xorm.NewEngine(driver, connection)
	if err != nil {
		return nil, err
	}
	db.ShowSQL(true)
	return db, nil
}

type Validator struct{}

func (v *Validator) Validate(i interface{}) error {
	_, err := govalidator.ValidateStruct(i)
	return err
}
