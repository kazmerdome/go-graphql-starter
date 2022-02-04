package provider_test

import (
	"testing"

	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/observe/logger"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider"

	"github.com/stretchr/testify/assert"
)

func TestNewProviderConfig(t *testing.T) {
	c := config.NewConfig(config.MODE_GLOBALENV)
	l := logger.NewZapLogger()
	s := provider.NewProviderConfig(l, c)

	assert.NotNil(t, s)
	assert.Equal(t, s.Config, c)
	assert.Equal(t, s.Logger, l)
}
