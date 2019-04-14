package mongo

import (
	"github.com/Nerzal/CardsOfBinokee-Server/pkg/core"
	"github.com/Nerzal/CardsOfBinokee-Server/pkg/repository"
	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

const collectionName = "cards"

// NewCardRepository creates a new instance of repository.CardRepository
func NewCardRepository() (repository.CardRepository, error) {
	dialInfo := getDialInfo()
	connection, err := bongo.Connect(&bongo.Config{DialInfo: dialInfo})
	return &cardRepository{connection}, err
}

type cardRepository struct {
	connection *bongo.Connection
}

func (repo *cardRepository) GetCards() []core.Card {
	var result []core.Card
	results := repo.connection.Collection(collectionName).Find(bson.M{})
	var card core.Card
	for results.Next(&card) {
		result = append(result, card)
	}
	return result
}

func (repo *cardRepository) SaveCards(cards []core.Card) error {
	collection := repo.connection.Collection(collectionName)
	collection.Delete(bson.M{})
	
	for _, card := range cards {
		err := collection.Save(&card)
		if err != nil {
			return err
		}
	}

	return nil
}
