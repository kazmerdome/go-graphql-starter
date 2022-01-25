package connector

import "github.com/kazmerdome/go-graphql-starter/pkg/auth/authorization/guard"

type GatewayGuards struct {
	LicenceGuard guard.LicenceGuard
}
