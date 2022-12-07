package foo

type Foo struct {
	ID int `json:"id"`
}

type FooResponse struct {
	Data []Foo
}
