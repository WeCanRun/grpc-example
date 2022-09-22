package auth

import (
	"context"
	"grpc-example/global/constants"
)

type Auth struct {
	AppKey    string
	AppSecret string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		constants.AppKey:    a.AppKey,
		constants.AppSecret: a.AppSecret,
	}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
	return false
}

func GetAppKey() string {
	return constants.AppKey
}

func GetAppSecret() string {
	return constants.AppSecret
}
