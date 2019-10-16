package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

var cardJSON = `{"object":"card","id":"25f2e4d0-effd-4e83-b7aa-1a0d8f120951","oracle_id":"713332c1-5bd8-400f-bfff-c1ca0697a043","name":"Opt","lang":"en","released_at":"2018-04-27","layout":"normal","mana_cost":"{U}","cmc":1,"type_line":"Instant","oracle_text":"Scry 1. (Look at the top card of your library. You may put that card on the bottom of your library.)\nDraw a card.","colors":["U"],"color_identity":["U"],"legalities":{"standard":"legal","future":"legal","historic":"legal","modern":"legal","legacy":"legal","pauper":"legal","vintage":"legal","penny":"not_legal","commander":"legal","brawl":"legal","duel":"legal","oldschool":"not_legal"},"games":["arena","mtgo","paper"],"reserved":false,"foil":true,"nonfoil":true,"oversized":false,"promo":false,"reprint":true,"variation":false,"set":"dom","set_name":"Dominaria","set_type":"expansion","set_uri":"https://api.scryfall.com/sets/be1daba3-51c9-4e7e-9212-36e68addc26c","set_search_uri":"https://api.scryfall.com/cards/search?order=set\u0026q=e%3Adom\u0026unique=prints","scryfall_set_uri":"https://scryfall.com/sets/dom?utm_source=api","rulings_uri":"https://api.scryfall.com/cards/25f2e4d0-effd-4e83-b7aa-1a0d8f120951/rulings","prints_search_uri":"https://api.scryfall.com/cards/search?order=released\u0026q=oracleid%3A713332c1-5bd8-400f-bfff-c1ca0697a043\u0026unique=prints","collector_number":"60","digital":false,"rarity":"common","flavor_text":"The crystal pulsed with the power of Teferi's planeswalker spark. Had Jhoira given him a blessing or a curse?","card_back_id":"0aeebaf5-8c7d-4636-9e82-8c27447861f7","artist":"Tyler Jacobson","artist_ids":["522af130-8db4-4b4b-950c-6e2b246339cf"],"illustration_id":"8be8d8ae-97c1-4da1-bc21-0aeb6c7082e9","border_color":"black","frame":"2015","full_art":false,"textless":false,"booster":true,"story_spotlight":false,"edhrec_rank":445}`

func TestGetCardByName(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(cardJSON))
	rec := httptest.NewRecorder()

	testContext := e.NewContext(req, rec)
	testContext.SetPath("/:card")
	testContext.SetParamNames("card")
	testContext.SetParamValues("opt")

	cardObj := Card{}

	_ = getCardByName(testContext)

	_ = json.Unmarshal([]byte(rec.Body.String()), &cardObj)
	if cardObj.Name != "Opt" {
		t.Errorf("%v is not Opt", string(cardObj.Name))
	} else {
		t.Logf("%v is Opt!", string(cardObj.Name))
	}
}

func TestGetCardBySetNum(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(cardJSON))
	rec := httptest.NewRecorder()

	testContext := e.NewContext(req, rec)
	testContext.SetPath("/:set/:num")
	testContext.SetParamNames("set", "num")
	testContext.SetParamValues("dom", "60")

	cardObj := Card{}

	_ = getCardBySetNum(testContext)

	_ = json.Unmarshal([]byte(rec.Body.String()), &cardObj)

	if cardObj.Name != "Opt" {
		t.Errorf("%v is not Opt", string(cardObj.Name))
	} else {
		t.Logf("%v is Opt!", string(cardObj.Name))
	}
	if cardObj.SetName != "Dominaria" {
		t.Errorf("%v is not Dominaria!", cardObj.SetName)
	} else {
		t.Logf("%v is Dominaria!", cardObj.SetName)
	}
}
