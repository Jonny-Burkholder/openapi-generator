package repo

import (
	"fmt"
	"test/openapi-generator/internal/bar"
)

type BarRepo struct{}

func (b *BarRepo) Find(id int) *bar.Bar {
	return &bar.Bar{ID: id}
}

func (b *BarRepo) FindAll() *bar.BarResponse {
	return &bar.BarResponse{
		Data: []bar.Bar{{ID: 1}, {ID: 2}},
	}
}

func (b *BarRepo) Create() error {
	fmt.Println("This is just an example repo, it doesn't actually create anything")
	return nil
}

func (b *BarRepo) Update() error {
	fmt.Println("Successfully updated")
	return nil
}

func (b *BarRepo) Delete() error {
	fmt.Println("Successfully \"deleted\" the fake object")
	return nil
}
