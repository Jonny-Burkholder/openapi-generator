package bar

type BarRepo interface {
	Find(id int) *Bar
	FindAll() *BarResponse
}
