package ladybug

import "github.com/ytorgon/ozon-route256-tg-bot/internal/model/vibrant"

type LadybugService interface {
	Describe(ladybugID uint64) (*vibrant.Ladybug, error)
	List(cursor uint64, limit uint64) ([]vibrant.Ladybug, error)
	Create(vibrant.Ladybug) (uint64, error)
	Update(ladybugID uint64, ladybug vibrant.Ladybug) error
	Remove(ladybugID uint64) (bool, error)
}

type DummyLadybugService struct{}

func NewDummyLadybugService() *DummyLadybugService {
	return &DummyLadybugService{}
}
