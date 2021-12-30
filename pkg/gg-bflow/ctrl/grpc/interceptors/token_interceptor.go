package interceptors

import (
	"context"
	"encoding/base64"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ctrl/grpc/errors"
	_grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func GetUnaryValidateTokenInterceptor(apiKey string, authorizedClients []string, whitelistMethod ...string) _grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *_grpc.UnaryServerInfo, handler _grpc.UnaryHandler) (resp interface{}, err error) {
		// continue if whitelisted
		if contains(whitelistMethod, info.FullMethod) {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.ErrGRPCUnauthenticated
		}

		a := md["authorization"]
		if len(a) < 1 {
			return nil, errors.ErrGRPCUnauthenticated
		}

		token := strings.TrimPrefix(a[0], "Basic ")
		b, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			return nil, errors.ErrMalformedToken
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			return nil, errors.ErrMalformedToken
		}

		if !contains(authorizedClients, pair[0]) {
			return nil, errors.ErrUnauthorizedToken
		}

		if pair[1] != apiKey {
			return nil, errors.ErrUnauthorizedToken
		}

		return handler(ctx, req)
	}
}

func GetStreamValidateTokenInterceptor(apiKey string, authorizedClients []string, whitelistMethod ...string) _grpc.StreamServerInterceptor {
	return func(srv interface{}, ss _grpc.ServerStream, info *_grpc.StreamServerInfo, handler _grpc.StreamHandler) error {
		// continue if whitelisted
		if contains(whitelistMethod, info.FullMethod) {
			return handler(srv, ss)
		}

		md, ok := metadata.FromIncomingContext(ss.Context())
		if !ok {
			return errors.ErrGRPCUnauthenticated
		}

		a := md["authorization"]
		if len(a) < 1 {
			return errors.ErrGRPCUnauthenticated
		}

		token := strings.TrimPrefix(a[0], "Basic ")
		b, err := base64.StdEncoding.DecodeString(token)
		if err != nil {
			return errors.ErrMalformedToken
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			return errors.ErrMalformedToken
		}

		if !contains(authorizedClients, pair[0]) {
			return errors.ErrUnauthorizedToken
		}

		if pair[1] != apiKey {
			return errors.ErrUnauthorizedToken
		}

		return handler(srv, ss)
	}
}
