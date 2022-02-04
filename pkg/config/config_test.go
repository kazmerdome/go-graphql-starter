package config_test

import (
	"os"
	"testing"

	"github.com/kazmerdome/go-graphql-starter/pkg/config"

	"github.com/stretchr/testify/assert"
)

const (
	k = "TEST_KEY"
	v = "TEST_VALUE"
)

func TestGet(t *testing.T) {
	s := config.NewConfig(config.MODE_GLOBALENV)
	os.Setenv(k, v)
	defer os.Unsetenv(k)

	// case New
	tv := s.Get(k)
	assert.Equal(t, tv, v)

	// case Empty
	tv2 := s.Get("XXX")
	assert.Equal(t, tv2, "")

	// case Cached
	tv3 := s.Get(k)
	assert.Equal(t, tv3, v)
}

func TestSet(t *testing.T) {
	s := config.NewConfig(config.MODE_GLOBALENV)

	// case New
	err := s.Set(k, v)
	tv := s.Get(k)
	assert.Equal(t, err, nil)
	assert.Equal(t, tv, v)

	// case Duplicated
	err = s.Set(k, v)
	assert.NotEqual(t, err, nil)
}
