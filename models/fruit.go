package models

import (
	"context"

	"github.com/go-xorm/xorm"

	"sample-go-api/factory"
)

type Fruit struct {
	Id        int64  `json:"id" xorm:"int64 notnull autoincr pk 'id'"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Price     int64  `json:"price"`
	StoreCode string `json:"storeCode"`
	CreatedAt string `json:"createdAt" xorm:"created"`
	UpdatedAt string `json:"updatedAt" xorm:"updated"`
}

func (d *Fruit) Create(ctx context.Context) (affectedRow int64, err error) {
	affectedRow, err = factory.DB(ctx).Insert(d)
	return
}
func (Fruit) GetById(ctx context.Context, id int64) (has bool, fruit *Fruit, err error) {
	fruit = &Fruit{}
	has, err = factory.DB(ctx).Where("id=?", id).Get(fruit)
	return
}
func (Fruit) GetAll(ctx context.Context, sortby, order []string, offset, limit int) (totalCount int64, items []Fruit, err error) {
	queryBuilder := func() *xorm.Session {
		q := factory.DB(ctx)
		if err := setSortOrder(q, sortby, order); err != nil {
			factory.Logger(ctx).Error(err)
		}
		return q
	}
	q := *queryBuilder()

	errc := make(chan error)
	go func(qNew xorm.Session) {
		v, err := qNew.Count(&Fruit{})
		if err != nil {
			errc <- err
			return
		}
		totalCount = v
		errc <- nil

	}(q)

	go func(qNew xorm.Session) {
		if err := qNew.Limit(limit, offset).Find(&items); err != nil {
			errc <- err
			return
		}
		errc <- nil
	}(q)

	if err := <-errc; err != nil {
		return 0, nil, err
	}
	if err := <-errc; err != nil {
		return 0, nil, err
	}
	return
}
func (d *Fruit) Update(ctx context.Context, id int64) (affectedRow int64, err error) {
	affectedRow, err = factory.DB(ctx).Where("id=?", id).Update(d)
	return
}

func (Fruit) Delete(ctx context.Context, id int64) (affectedRow int64, err error) {
	affectedRow, err = factory.DB(ctx).Where("id=?", id).Delete(&Fruit{})
	return
}
