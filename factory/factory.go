package factory

import (
	"context"

	"github.com/go-xorm/xorm"
	"github.com/pangpanglabs/goutils/behaviorlog"
	"github.com/pangpanglabs/goutils/echomiddleware"
	"github.com/sirupsen/logrus"
)

func DB(ctx context.Context) *xorm.Session {
	v := ctx.Value(echomiddleware.ContextDBName)
	if v == nil {
		panic("DB is not exist")
	}
	if db, ok := v.(*xorm.Session); ok {
		return db
	}
	if db, ok := v.(*xorm.Engine); ok {
		return db.NewSession()
	}
	panic("DB is not exist")
}

func BehaviorLogger(ctx context.Context) *behaviorlog.LogContext {
	v := ctx.Value(behaviorlog.LogContextName)
	if logger, ok := v.(*behaviorlog.LogContext); ok {
		return logger.Clone()
	}
	return behaviorlog.NewNopContext()
}

func Logger(ctx context.Context) *logrus.Entry {
	v := ctx.Value(behaviorlog.LogContextName)
	if v == nil {
		return logrus.WithFields(logrus.Fields{})
	}
	if logger, ok := v.(*logrus.Entry); ok {
		return logger
	}
	return logrus.WithFields(logrus.Fields{})
}

const (
	ContextConfigName = "sample-go-api-config"
)

func Config(ctx context.Context, key string) interface{} {
	return config(ctx, key)
}

func ConfigString(ctx context.Context, key string) string {
	value, has := config(ctx, key).(string)
	if has {
		return value
	}
	return ""
}

func config(ctx context.Context, key string) interface{} {
	v := ctx.Value(ContextConfigName)
	if v == nil {
		panic("config context is not exist")
	}
	if customConfig, ok := v.(*map[string]interface{}); ok {
		if value, has := (*customConfig)[key]; has {
			return value
		}
	}
	panic("config context is not exist")
}
