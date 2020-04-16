package main

import (
	"testing"

	"github.com/linkedin/goavro"
	"github.com/stretchr/testify/assert"
)

func TestAvroLinkedinDecode(t *testing.T) {
	codec, err := goavro.NewCodec(ArvoSuperheroSchema)
	assert.NoError(t, err)

	superhero, _, err := codec.NativeFromBinary(ArvoSuperheroBytes)
	assert.NoError(t, err)
	assert.Equal(t, *SuperheroLinkedinFixture, superhero)
}

func BenchmarkAvroLinkedinDecode(b *testing.B) {
	codec, _ := goavro.NewCodec(ArvoSuperheroSchema)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = codec.NativeFromBinary(ArvoSuperheroBytes)
	}
}

func TestAvroLinkedinEncode(t *testing.T) {
	codec, err := goavro.NewCodec(ArvoSuperheroSchema)
	assert.NoError(t, err)

	bytes, err := codec.BinaryFromNative(nil, *SuperheroLinkedinFixture)
	assert.NoError(t, err)
	assert.Equal(t, ArvoSuperheroSimpleBytes, bytes)
}

func BenchmarkAvroLinkedinEncode(b *testing.B) {
	codec, _ := goavro.NewCodec(ArvoSuperheroSchema)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = codec.BinaryFromNative(nil, *SuperheroLinkedinFixture)
	}
}
