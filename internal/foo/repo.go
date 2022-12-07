package foo

type FooRepo interface {
	Find(id int) *Foo
	FindAll() *FooResponse
}
