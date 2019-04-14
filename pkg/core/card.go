package core

// Card represents a card
type Card struct {
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
