package main

import (
	"testing"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/stretchr/testify/assert"
)

func TestJSONPquernaFfjsonDecode(t *testing.T) {
	superhero := Superhero{}

	err := ffjson.Unmarshal(SuperheroJSONBytes, &superhero)
	assert.NoError(t, err)
	assert.Equal(t, *SuperheroFixture, superhero)
}

func BenchmarkJSONPquernaFfjsonDecode(b *testing.B) {
	superhero := Superhero{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ffjson.Unmarshal(SuperheroJSONBytes, &superhero)
	}
}

func TestJSONPquernaFfjsonEncode(t *testing.T) {
	bytes, err := ffjson.Marshal(*SuperheroFixture)
	assert.NoError(t, err)
	assert.Equal(t, SuperheroJSONBytes, bytes)
}

func BenchmarkJSONPquernaFfjsonEncode(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ffjson.Marshal(*SuperheroFixture)
	}
}
