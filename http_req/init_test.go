package http_req

import (
	"context"
	"os"
	"sample-go-api/factory"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ctx context.Context
)

func init() {
	configMap := map[string]interface{}{
		"sample_url": os.Getenv("Sample_Url"),
	}
	ctx = context.WithValue(context.Background(), factory.ContextConfigName, &configMap)
}
