package foo

type FooRepo interface {
	Find(id int) *Foo
	FindAll() *FooResponse
	Create() error
	Update() error
	Delete() error
}
