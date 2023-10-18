package watcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrettyPrintFloat(t *testing.T) {
	assert.Equal(t, "100", prettyPrintFloat(100))
	assert.Equal(t, "99.42", prettyPrintFloat(99.42098176))
}

func TestPrettyPrintAccountID(t *testing.T) {
	assert.Equal(t, "foo", prettyPrintAccountID("foo.poolv1.near"))
	assert.Equal(t, "foo", prettyPrintAccountID("foo.pool.near"))
	assert.Equal(t, "foo", prettyPrintAccountID("foo.pool.f863973.m0"))
}
