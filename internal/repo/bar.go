package repo

import (
	"test/openapi-generator/internal/bar"
)

type BarRepo struct{}

func (B *BarRepo) Find(id int) *bar.Bar {
	return &bar.Bar{ID: id}
}

func (B *BarRepo) FindAll() *bar.BarResponse {
	return &bar.BarResponse{
		Data: []bar.Bar{{ID: 1}, {ID: 2}},
	}
}
