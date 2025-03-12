package token

import "context"

const (
	AppName = "token"
)

type Service interface {
	RPCServer
	RevolkToken(context.Context, *RevolkTokenRequest) (*Token, error)
}
