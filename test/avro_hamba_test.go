package main

import (
	"testing"

	"github.com/hamba/avro"
	"github.com/stretchr/testify/assert"
)

func TestAvroHambaDecode(t *testing.T) {
	schema := avro.MustParse(ArvoSuperheroSchema)
	superhero := Superhero{}

	err := avro.Unmarshal(schema, ArvoSuperheroBytes, &superhero)
	assert.NoError(t, err)
	assert.Equal(t, *SuperheroFixture, superhero)
}

func BenchmarkAvroHambaDecode(b *testing.B) {
	schema := avro.MustParse(ArvoSuperheroSchema)
	superhero := Superhero{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = avro.Unmarshal(schema, ArvoSuperheroBytes, &superhero)
	}
}

func TestAvroHambaEncode(t *testing.T) {
	schema := avro.MustParse(ArvoSuperheroSchema)

	bytes, err := avro.Marshal(schema, *SuperheroFixture)
	assert.NoError(t, err)
	assert.Equal(t, ArvoSuperheroBytes, bytes)
}

func BenchmarkAvroHambaEncode(b *testing.B) {
	schema := avro.MustParse(ArvoSuperheroSchema)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = avro.Marshal(schema, *SuperheroFixture)
	}
}
