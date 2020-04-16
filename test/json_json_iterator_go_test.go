package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	jsoniter "github.com/json-iterator/go"
)

func TestJSONIteratorGoDecode(t *testing.T) {
	superhero := Superhero{}

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(SuperheroJSONBytes, &superhero)
	assert.NoError(t, err)
	assert.Equal(t, *SuperheroFixture, superhero)
}

func BenchmarkJSONIteratorGoDecode(b *testing.B) {
	superhero := Superhero{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(SuperheroJSONBytes, &superhero)
	}
}

func TestJSONIteratorGoEncode(t *testing.T) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	bytes, err := json.Marshal(*SuperheroFixture)
	assert.NoError(t, err)
	assert.Equal(t, SuperheroJSONBytes, bytes)
}

func BenchmarkJSONIteratorGoEncode(b *testing.B) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(*SuperheroFixture)
	}
}
