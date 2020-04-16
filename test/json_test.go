package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONDecode(t *testing.T) {
	superhero := Superhero{}

	err := json.Unmarshal(SuperheroJSONBytes, &superhero)
	assert.NoError(t, err)
	assert.Equal(t, *SuperheroFixture, superhero)
}

func BenchmarkJSONDecode(b *testing.B) {
	superhero := Superhero{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(SuperheroJSONBytes, &superhero)
	}
}

func TestJSONEncode(t *testing.T) {
	bytes, err := json.Marshal(*SuperheroFixture)
	assert.NoError(t, err)
	assert.Equal(t, SuperheroJSONBytes, bytes)
}

func BenchmarkJSONEncode(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(*SuperheroFixture)
	}
}
