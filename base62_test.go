package base62

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	id = 23029975486464
	s = "6xsfs9JC"
)

func TestBase62(t *testing.T) {
	t.Run("encode", func(t *testing.T) {
		assert.Equal(t, s, Encode(id))
	})

	t.Run("decode", func(t *testing.T) {
		assert.Equal(t, id, Decode(s))
	})
}