package models

import (
	"fmt"
	"testing"
)

func TestFruitCreate(t *testing.T) {
	f := &Fruit{
		Code: "123",
	}
	affectedRow, err := f.Create(ctx)
	fmt.Println(affectedRow, err, f)
}

func TestFruitUpdate(t *testing.T) {
	f := &Fruit{
		Code: "222",
	}
	affectedRow, err := f.Update(ctx, 1)
	fmt.Println(affectedRow, err)
}

func TestFruitDelete(t *testing.T) {
	affectedRow, err := Fruit{}.Delete(ctx, 2)
	fmt.Println(affectedRow, err)
}

func TestFruitGetAll(t *testing.T) {
	total, items, err := Fruit{}.GetAll(ctx, nil, nil, 0, 2)
	fmt.Println(total, items, err)
}
func TestFruitGetById(t *testing.T) {
	has, v, err := Fruit{}.GetById(ctx, 1)
	fmt.Println(has, v, err)
}
