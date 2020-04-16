package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/snappy"
	"github.com/hamba/avro"
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

func main()  {
	payload, err := ioutil.ReadFile("test/fixtures/superhero.avsc")
	if err != nil {
		log.Fatal(err)
	}
	schema := avro.MustParse(string(payload))

	superhero := &Superhero{}
	payload, err = ioutil.ReadFile("test/fixtures/superhero.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(payload, &superhero)
	if err != nil {
		log.Fatal(err)
	}
	
	bytes, err := avro.Marshal(schema, &superhero)
	if err != nil {
		log.Fatal(err)
	}

	bytes = snappy.Encode(nil, bytes)

	path := "test/fixtures/avro-hamba-snappy-superhero.bin"

	err = ioutil.WriteFile(path, bytes, 0755)
	if err != nil {
		log.Fatal(err)
	}

	payload, err = ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	payload, err = snappy.Decode(nil, payload)
	if err != nil {
		log.Fatal(err)
	}

	var result Superhero
	err = avro.Unmarshal(schema, payload, &result)
	if err != nil {
		log.Fatal(err)
	}

	if superhero.ID != result.ID {
		log.Fatal("Wrong result")
	}

	fmt.Println(result.ID)
}