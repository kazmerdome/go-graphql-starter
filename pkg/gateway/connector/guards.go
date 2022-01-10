package connector

import "github.com/kazmerdome/go-graphql-starter/pkg/auth/authorization/guards"

type GatewayGuards struct {
	LicenceGuard guards.LicenceGuard
}
