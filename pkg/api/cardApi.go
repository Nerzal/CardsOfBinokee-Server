package api

import (
	"net/http"

	"github.com/Nerzal/CardsOfBinokee-Server/pkg/card"
	"github.com/Nerzal/CardsOfBinokee-Server/pkg/core"
	"github.com/labstack/echo"
)

type CardAPI interface {
	PostCards(c echo.Context) error
	GetCards(c echo.Context) error
}

type cardAPI struct {
	handler card.Handler
}

func NewCardAPI(handler card.Handler) CardAPI {
	return &cardAPI{
		handler: handler,
	}
}

func (cardAPI *cardAPI) PostCards(c echo.Context) error {
	var request []core.Card
	err := c.Bind(&request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "U failed!")
	}

	err = cardAPI.handler.PostCards(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Oh noes :/")
	}

	return c.JSON(http.StatusOK, "ok")
}

func (cardAPI *cardAPI) GetCards(c echo.Context) error {
	result := cardAPI.handler.GetCards()
	return c.JSON(http.StatusOK, result)
}
