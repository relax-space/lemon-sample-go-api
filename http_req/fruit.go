package http_req

import (
	"context"
	"fmt"
	"goutils/httpreq"
	"net/http"
	"sample-go-api/factory"
	"sample-go-api/models"

	"github.com/relax-space/go-kit/model"
)

type HttpReq struct {
}

func (HttpReq) GetAllFruit(ctx context.Context) (status int, result *model.Result, err error) {
	result = &model.Result{}
	url := fmt.Sprintf("%v/fruits", factory.ConfigString(ctx, "sample_url"))
	status, err = httpreq.New(http.MethodGet, url, nil).Call(result)
	return
}

func (HttpReq) PostFruit(ctx context.Context, fruit *models.Fruit) (status int, result *model.Result, err error) {
	result = &model.Result{}
	url := fmt.Sprintf("%v/fruits", factory.ConfigString(ctx, "sample_url"))
	status, err = httpreq.New(http.MethodPost, url, fruit).Call(result)
	return
}
