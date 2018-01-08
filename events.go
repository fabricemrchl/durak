package main

import (
	"encoding/json"
	"reflect"
)

// JSONEvent represents a message to clients with some event.
type JSONEvent struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

// GameInList contains short info about game.
type GameInList struct {
	Name string `json:"name"`
	Id   uint64 `json:"id"`
}

// ClientJoinedEvent contains info for the just connected client.
type ClientJoinedEvent struct {
	Id         uint64        `json:"id"`
	Nickname   string        `json:"nickname"`
	ClientsNum int           `json:"clients_num"`
	Games      []*GameInList `json:"games"`
}

// ClientBroadCastJoinedEvent contains info for other clients when a new client was connected.
type ClientBroadCastJoinedEvent struct {
	Id         uint64 `json:"id"`
	Nickname   string `json:"nickname"`
	ClientsNum int    `json:"clients_num"`
}

// PlayersEvent contains list of players which were connected to a game.
type PlayersEvent struct {
	YourPlayerIndex int       `json:"your_player_index"`
	Players         []*Player `json:"players"`
}

// DealEvent contains info about game after the deal. It includes list of cards for each player.
type DealEvent struct {
	YourHand                      []*Card `json:"your_hand"`
	HandsSizes                    []int   `json:"hands_sizes"`
	PileSize                      int     `json:"pile_size"`
	TrumpSuit                     string  `json:"trump_suit"`
	TrumpCard                     *Card   `json:"trump_card"`
	TrumpCardIsInPile             bool    `json:"trump_card_is_in_pile"`
	TrumpCardIsOwnedByPlayerIndex int     `json:"trump_card_is_owned_by_player_index"`
}

// FirstAttackerEvent contains info who is the first attacker and why.
type FirstAttackerEvent struct {
	ReasonCard    *Card `json:"reason_card"`
	AttackerIndex int   `json:"attacker_index"`
}

func eventToJSON(e interface{}) ([]byte, error) {
	name := getNameOfStruct(e)
	jsonEvent := JSONEvent{Name: name, Data: e}
	return json.Marshal(jsonEvent)
}

func getNameOfStruct(s interface{}) string {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}
	return t.Name()
}
