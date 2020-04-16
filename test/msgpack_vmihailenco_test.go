package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v4"
)

func TestMsgpacVmihailencoDecode(t *testing.T) {
	superhero := Superhero{}

	err := msgpack.Unmarshal(MsgpacSuperheroBytes, &superhero)
	assert.NoError(t, err)
	assert.Equal(t, *SuperheroFixture, superhero)
}

func BenchmarkMsgpacVmihailencoDecode(b *testing.B) {
	superhero := Superhero{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = msgpack.Unmarshal(MsgpacSuperheroBytes, &superhero)
	}
}

func TestMsgpacVmihailencoEncode(t *testing.T) {
	bytes, err := msgpack.Marshal(*SuperheroFixture)
	assert.NoError(t, err)
	assert.Equal(t, MsgpacSuperheroBytes, bytes)
}

func BenchmarkMsgpacVmihailencoEncode(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = msgpack.Marshal(*SuperheroFixture)
	}
}
