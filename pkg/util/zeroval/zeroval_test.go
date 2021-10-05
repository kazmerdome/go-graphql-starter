package zeroval_test

import (
	"testing"

	"github.com/kazmerdome/go-graphql-starter/pkg/util/zeroval"

	"github.com/stretchr/testify/assert"
)

type testdata struct{}

func TestIsZeroVal(t *testing.T) {
	refType := testdata{}
	valType := ""

	assert.Zero(t, refType)
	assert.Zero(t, valType)

	assert.True(t, zeroval.IsZeroVal(refType))
	assert.True(t, zeroval.IsZeroVal(valType))
}

func TestIsNotZeroVal(t *testing.T) {
	refType := new(testdata)
	valType := "!"

	assert.NotZero(t, refType)
	assert.NotZero(t, valType)

	assert.False(t, zeroval.IsZeroVal(refType))
	assert.False(t, zeroval.IsZeroVal(valType))
}
