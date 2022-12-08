package repo

import (
	"fmt"
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

func (f *FooRepo) Create() error {
	fmt.Println("This is just an example repo, it doesn't actually create anything")
	return nil
}

func (f *FooRepo) Update() error {
	fmt.Println("Successfully updated")
	return nil
}

func (f *FooRepo) Delete() error {
	fmt.Println("Successfully \"deleted\" the fake object")
	return nil
}
