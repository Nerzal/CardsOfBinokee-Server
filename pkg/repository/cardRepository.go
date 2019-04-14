package repository

import "github.com/Nerzal/CardsOfBinokee-Server/pkg/core"

type CardRepository interface {
	GetCards() []core.Card
	SaveCards(cards []core.Card) error
}
