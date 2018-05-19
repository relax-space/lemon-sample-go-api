package http_req

import (
	"fmt"
	"net/http"
	"sample-go-api/models"
	"testing"

	"github.com/relax-space/go-kit/test"
)

func Test_GetAllFruit(t *testing.T) {
	status, fruits, err := HttpReq{}.GetAllFruit(ctx)
	test.Equals(t, http.StatusOK, status)
	test.Ok(t, err)
	fmt.Println(fruits)

}

func Test_PostFruit(t *testing.T) {
	fruit := &models.Fruit{
		Code:  "1",
		Color: "red",
	}
	status, fruits, err := HttpReq{}.PostFruit(ctx, fruit)
	test.Equals(t, http.StatusCreated, status)
	test.Ok(t, err)
	fmt.Println(fruits)
}
