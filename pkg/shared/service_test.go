package shared_test

import (
	"testing"

	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/logger"
	"github.com/kazmerdome/go-graphql-starter/pkg/shared"

	"github.com/stretchr/testify/assert"
)

func TestNewBaseService(t *testing.T) {
	c := config.NewConfigService(config.MODE_GLOBALENV)
	l := logger.NewZapLogger()
	s := shared.NewSharedService(l, c)

	assert.NotNil(t, s)
	assert.Equal(t, s.Config, c)
	assert.Equal(t, s.Logger, l)
}
