package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/avro.v0"
)

func TestAvroAvroGoAvroDecode(t *testing.T) {
	schema := avro.MustParseSchema(ArvoSuperheroSchema)
	reader := avro.NewSpecificDatumReader()
	reader.SetSchema(schema)
	superhero := Superhero{}

	decoder := avro.NewBinaryDecoder(ArvoSuperheroBytes)
	err := reader.Read(&superhero, decoder)
	assert.NoError(t, err)
	assert.Equal(t, *SuperheroFixture, superhero)
}

func BenchmarkAvroGoAvroDecode(b *testing.B) {
	schema := avro.MustParseSchema(ArvoSuperheroSchema)
	reader := avro.NewSpecificDatumReader()
	reader.SetSchema(schema)
	superhero := Superhero{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		decoder := avro.NewBinaryDecoder(ArvoSuperheroBytes)
		_ = reader.Read(&superhero, decoder)
	}
}

func TestAvroGoAvroEncode(t *testing.T) {
	schema := avro.MustParseSchema(ArvoSuperheroSchema)
	writer := avro.NewSpecificDatumWriter()
	writer.SetSchema(schema)
	buf := &bytes.Buffer{}
	encoder := avro.NewBinaryEncoder(buf)

	err := writer.Write(SuperheroFixture, encoder)
	assert.Equal(t, ArvoSuperheroSimpleBytes, buf.Bytes())
	assert.NoError(t, err)
}

func BenchmarkAvroGoAvroEncode(b *testing.B) {
	schema := avro.MustParseSchema(ArvoSuperheroSchema)
	writer := avro.NewSpecificDatumWriter()
	writer.SetSchema(schema)
	buf := &bytes.Buffer{}
	encoder := avro.NewBinaryEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = writer.Write(SuperheroFixture, encoder)
	}
}
