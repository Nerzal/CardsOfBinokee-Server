package core

import "github.com/go-bongo/bongo"

// Card represents a card
type Card struct {
	bongo.DocumentBase `json:"-" bson:",inline"`

	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	Up           int    `json:"up"`
	Right        int    `json:"right"`
	Down         int    `json:"down"`
	Left         int    `json:"left"`
	BackgroundID int    `json:"backgroundID"`
	CreatureID   int    `json:"creatureID"`
}

