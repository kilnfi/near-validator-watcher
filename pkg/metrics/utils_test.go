package metrics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolToFloat64(t *testing.T) {
	assert.Equal(t, float64(1), BoolToFloat64(true))
	assert.Equal(t, float64(0), BoolToFloat64(false))
}

func TestStringToFloat64(t *testing.T) {
	assert.NotEqual(t, float64(0), StringToFloat64("foobar"))
}
