package main

import (
	"testing"

	"github.com/golang/snappy"
	"github.com/hamba/avro"
	"github.com/stretchr/testify/assert"
)

func TestAvroHambaSnappyDecode(t *testing.T) {
	schema := avro.MustParse(ArvoSuperheroSchema)
	superhero := Superhero{}

	bytes, err := snappy.Decode(nil, ArvoSuperheroSnappyBytes)
	assert.NoError(t, err)
	err = avro.Unmarshal(schema, bytes, &superhero)
	assert.NoError(t, err)
	assert.Equal(t, *SuperheroFixture, superhero)
}

func BenchmarkAvroHambaSnappyDecode(b *testing.B) {
	schema := avro.MustParse(ArvoSuperheroSchema)
	superhero := Superhero{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes, _ := snappy.Decode(nil, ArvoSuperheroSnappyBytes)
		_ = avro.Unmarshal(schema, bytes, &superhero)
	}
}

func TestAvroHambaSnappyEncode(t *testing.T) {
	schema := avro.MustParse(ArvoSuperheroSchema)

	bytes, err := avro.Marshal(schema, *SuperheroFixture)
	result := snappy.Encode(nil, bytes)
	assert.NoError(t, err)
	assert.Equal(t, ArvoSuperheroSnappyBytes, result)
}

func BenchmarkAvroHambaSnappyEncode(b *testing.B) {
	schema := avro.MustParse(ArvoSuperheroSchema)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bytes, _ := avro.Marshal(schema, *SuperheroFixture)
		_ = snappy.Encode(nil, bytes)
	}
}
