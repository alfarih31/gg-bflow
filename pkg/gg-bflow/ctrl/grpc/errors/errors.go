package errors

import (
	"github.com/alfarih31/gg-bflow/configs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrGRPCUnauthenticated = status.Errorf(codes.Unauthenticated, "%s", "unauthenticated")
var ErrMalformedToken = status.Errorf(codes.Unauthenticated, "%s", "malformed authorization, expected: \"Basic username:password\"")
var ErrUnauthorizedToken = status.Errorf(codes.Unauthenticated, "%s", "unauthorized authorization")
var ErrNotFound = status.Errorf(codes.NotFound, "%s", "not found")
var ErrBufferLimitExceed = status.Errorf(codes.InvalidArgument, "%s: %d byte", "buffer size exceed limit", configs.GGBFlow.BufferSizeLimit)
