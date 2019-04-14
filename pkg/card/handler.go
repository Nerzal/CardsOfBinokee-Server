package card

import (
	"github.com/Nerzal/CardsOfBinokee-Server/pkg/core"
	"github.com/Nerzal/CardsOfBinokee-Server/pkg/repository"
)

// Handler is a handler for cards
type Handler interface {
	GetCards() []core.Card
	PostCards(cards []core.Card) error
}

type handler struct {
	repo repository.CardRepository
}

// NewHandler returns a new instance of Handler
func NewHandler(repository repository.CardRepository) Handler {
	return &handler{
		repo: repository,
	}
}

func (handler *handler) GetCards() []core.Card {
	return handler.repo.GetCards()
}

func (handler *handler) PostCards(cards []core.Card) error {
	return handler.repo.SaveCards(cards)
}
