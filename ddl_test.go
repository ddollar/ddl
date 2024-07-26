package ddl_test

import (
	"testing"

	"github.com/ddollar/ddl"
	"github.com/stretchr/testify/assert"
)

type ifCase struct {
	condition bool
	thenValue any
	elseValue any
}

func TestIf(t *testing.T) {
	assert.Equal(t, 1, ddl.If(true, 1, 2))
	assert.Equal(t, 2, ddl.If(false, 1, 2))
}
