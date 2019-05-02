package api

import (
	"io/ioutil"
	"net/http"

	"github.com/Nerzal/CardsOfBinokee-Server/pkg/card"
	"github.com/Nerzal/CardsOfBinokee-Server/pkg/core"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
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
	var request CardsResponse
	err := c.Bind(&request)

	if err != nil {
		log.Error("Initial Error: ", err)
		requestBodyBytes, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error(err)
		}
		log.Info(string(requestBodyBytes))
		log.Error("Failed to deserialize request body: ", err)
		return c.JSON(http.StatusBadRequest, "U failed!")
	}

	if len(request.Items) == 0 {
		requestBodyBytes, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error(err)
		}
		log.Info(string(requestBodyBytes))
		log.Error("Failed to save cards. Empty array not allowed!")
		return c.JSON(http.StatusBadRequest, "Empty Array!")
	}

	err = cardAPI.handler.PostCards(request.Items)
	if err != nil {
		log.Error("Failed to handle PostCards request: ", err)
		return c.JSON(http.StatusInternalServerError, "Oh noes :/")
	}

	log.Info("Successfully updated card database!")

	return c.JSON(http.StatusOK, "ok")
}

func (cardAPI *cardAPI) GetCards(c echo.Context) error {
	result := cardAPI.handler.GetCards()
	response := CardsResponse{Items: result}
	return c.JSON(http.StatusOK, response)
}

type CardsResponse struct {
	Items []core.Card `json:"items"`
}
