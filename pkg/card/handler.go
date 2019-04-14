package card

import "github.com/Nerzal/CardsOfBinokee-Server/pkg/core"

// Handler is a handler for cards
type Handler interface {
	GetCards() ([]core.Card, error)
	PostCards(cards []core.Card) error
}

type handler struct {
}

// NewHandler returns a new instance of Handler
func NewHandler() Handler {
	return &handler{}
}

func (handler *handler) GetCards() ([]core.Card, error) {
	var result []core.Card
	result = append(result, core.Card{
		ID:           1,
		Name:         "Meister Metin",
		Description:  "Ist echt richtig stark!",
		Up:           10,
		Right:        9,
		Down:         8,
		Left:         10,
		BackgroundID: 1337,
		CreatureID:   42,
	})

	result = append(result, core.Card{
		ID:           2,
		Name:         "Meister Nerzal",
		Description:  "Ist echt viel stärker als Meister Metin!",
		Up:           10,
		Right:        10,
		Down:         10,
		Left:         10,
		BackgroundID: 1337,
		CreatureID:   42,
	})

	result = append(result, core.Card{
		ID:           3,
		Name:         "Meister Daniel",
		Description:  "Ist echt nicht so stark wie Meister Nerzal!",
		Up:           9,
		Right:        9,
		Down:         9,
		Left:         9,
		BackgroundID: 1337,
		CreatureID:   42,
	})

	return result, nil
}

func (handler *handler) PostCards(cards []core.Card) error {
	return nil
}
