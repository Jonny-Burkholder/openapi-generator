package foo

type FooService struct {
	FooRepo FooRepo
}

func NewFooService(FooRepo FooRepo) *FooService {
	return &FooService{FooRepo}
}

func (f *FooService) Find(id int) *Foo {
	return f.FooRepo.Find(id)
}

func (f *FooService) FindAll() *FooResponse {
	return f.FooRepo.FindAll()
}
