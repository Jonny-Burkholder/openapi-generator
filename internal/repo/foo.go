package repo

import (
	"test/openapi-generator/internal/foo"
)

type FooRepo struct{}

func (F *FooRepo) Find(id int) *foo.Foo {
	return &foo.Foo{ID: id}
}

func (F *FooRepo) FindAll() *foo.FooResponse {
	return &foo.FooResponse{
		Data: []foo.Foo{{ID: 1}, {ID: 2}},
	}
}
