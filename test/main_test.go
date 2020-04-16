package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

type Superhero struct {
	ID            int32         `avro:"id"`
	AffiliationID int32         `avro:"affiliation_id"`
	Name          string        `avro:"name"`
	Life          float32       `avro:"life"`
	Energy        float32       `avro:"energy"`
	Powers        []*Superpower `avro:"powers"`
}

type Superpower struct {
	ID      int32   `avro:"id"`
	Name    string  `avro:"name"`
	Damage  float32 `avro:"damage"`
	Energy  float32 `avro:"energy"`
	Passive bool    `avro:"passive"`
}


var SuperheroLinkedinFixture *map[string]interface{}

var SuperheroFixture *Superhero

var SuperheroJSONBytes []byte

var ArvoSuperheroSchema string

var ArvoSuperheroBytes []byte

var ArvoSuperheroSnappyBytes []byte

// The spec for Avro make blocks (arrays) size optional,
// this payload does not include it.
var ArvoSuperheroSimpleBytes []byte

var MsgpacSuperheroBytes []byte

func TestMain(m *testing.M) {
	SuperheroLinkedinFixture = &map[string]interface{}{
		"id":             int32(234765),
		"affiliation_id": int32(9867),
		"name":           "Wolverine",
		"life":           float32(85.25),
		"energy":         float32(32.75),
		"powers": []interface{}{
			map[string]interface{}{"id": int32(2345), "name": "Bone Claws", "damage": float32(5), "energy": float32(1.15), "passive": false},
			map[string]interface{}{"id": int32(2346), "name": "Regeneration", "damage": float32(-2), "energy": float32(0.55), "passive": true},
			map[string]interface{}{"id": int32(2347), "name": "Adamant skeleton", "damage": float32(-10), "energy": float32(0), "passive": true},
		},
	}

	var payload []byte
	var err error

	payload, err = ioutil.ReadFile("fixtures/superhero.json")
	if err != nil {
		log.Fatal(err)
	}
	SuperheroJSONBytes = payload

	SuperheroFixture = &Superhero{}
	err = json.Unmarshal(SuperheroJSONBytes, &SuperheroFixture)
	if err != nil {
		log.Fatal(err)
	}
	
	payload, err = ioutil.ReadFile("fixtures/superhero.avsc")
	if err != nil {
		log.Fatal(err)
	}
	ArvoSuperheroSchema = string(payload)

	payload, err = ioutil.ReadFile("fixtures/avro-hamba-superhero.bin")
	if err != nil {
		log.Fatal(err)
	}
	ArvoSuperheroBytes = payload

	payload, err = ioutil.ReadFile("fixtures/avro-hamba-snappy-superhero.bin")
	if err != nil {
		log.Fatal(err)
	}
	ArvoSuperheroSnappyBytes = payload

	payload, err = ioutil.ReadFile("fixtures/avro-go-arvo-simple-superhero.bin")
	if err != nil {
		log.Fatal(err)
	}
	ArvoSuperheroSimpleBytes = payload

	payload, err = ioutil.ReadFile("fixtures/msgpac-vmihailenco-superhero.bin")
	if err != nil {
		log.Fatal(err)
	}
	MsgpacSuperheroBytes = payload

	os.Exit(m.Run())
}
