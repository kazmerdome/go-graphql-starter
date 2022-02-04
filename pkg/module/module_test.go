package module_test

import (
	"testing"

	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/observe/logger"

	"github.com/stretchr/testify/assert"
)

func TestNewModuleConfig(t *testing.T) {
	c := config.NewConfig(config.MODE_GLOBALENV)
	l := logger.NewZapLogger()
	s := module.NewModuleConfig(l, c)

	assert.NotNil(t, s)
	assert.Equal(t, s.GetProviderConfig().Config, c)
	assert.Equal(t, s.GetProviderConfig().Logger, l)
}
