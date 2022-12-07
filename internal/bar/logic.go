package bar

type BarService struct {
	BarRepo BarRepo
}

func NewBarService(BarRepo BarRepo) *BarService {
	return &BarService{BarRepo}
}

func (b *BarService) Find(id int) *Bar {
	return b.BarRepo.Find(id)
}

func (b *BarService) FindAll() *BarResponse {
	return b.BarRepo.FindAll()
}
