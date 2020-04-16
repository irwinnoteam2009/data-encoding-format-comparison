package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/avro.v0"
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
	schema := avro.MustParseSchema(string(payload))

	superhero := &Superhero{}
	payload, err = ioutil.ReadFile("test/fixtures/superhero.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(payload, &superhero)
	if err != nil {
		log.Fatal(err)
	}

	writer := avro.NewSpecificDatumWriter()
	writer.SetSchema(schema)

	buf := &bytes.Buffer{}
	encoder := avro.NewBinaryEncoder(buf)

	err = writer.Write(superhero, encoder)
	if err != nil {
		log.Fatal(err)
	}

	path := "test/fixtures/avro-go-arvo-simple-superhero.bin"
	
	err = ioutil.WriteFile(path, buf.Bytes(), 0755)
	if err != nil {
		log.Fatal(err)
	}

	payload, err = ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var result Superhero
	decoder := avro.NewBinaryDecoder(payload)
	reader := avro.NewSpecificDatumReader()
	reader.SetSchema(schema)
	err = reader.Read(&result, decoder)
	if err != nil {
		log.Fatal(err)
	}

	if superhero.ID != result.ID {
		log.Fatal("Wrong result")
	}

	fmt.Println(result.ID)
}