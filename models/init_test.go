package models

import (
	"context"
	"os"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pangpanglabs/goutils/echomiddleware"
)

var ctx context.Context

func init() {
	runtime.GOMAXPROCS(1)
	xormEngine, err := xorm.NewEngine("mysql", os.Getenv("FRUIT_CONN"))
	if err != nil {
		panic(err)
	}
	xormEngine.ShowSQL(true)
	xormEngine.Sync(new(Fruit))

	ctx = context.WithValue(context.Background(), echomiddleware.ContextDBName, xormEngine.NewSession())
}
