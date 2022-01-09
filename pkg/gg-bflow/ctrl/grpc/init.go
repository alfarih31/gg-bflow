package grpc_ctrl

import (
	"fmt"
	"github.com/alfarih31/gg-bflow/api/grpc"
	"github.com/alfarih31/gg-bflow/configs"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ctrl/grpc/interceptors"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/logger"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/svc/buffer"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/svc/meta"
	_grpc "google.golang.org/grpc"
	"net"
	"time"
)

type gRPCCtrl struct {
	gg_bflow.UnimplementedGGBFlowServer
	st        time.Time
	bufferSvc buffer.BufferSvc
	metaSvc   meta.MetaSvc
}

var instance = &gRPCCtrl{
	bufferSvc: buffer.Svc,
	metaSvc:   meta.Svc,
}

var s *_grpc.Server

func Init() error {
	instance.st = time.Now()
	return instance.Init()
}

func Start() error {
	cfg := configs.GGBFlow

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.ServerName, cfg.Port))

	logger.Log.Info(fmt.Sprintf("gRPC Controller running: %s", lis.Addr()))
	err = s.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}

func (f *gRPCCtrl) Init() error {
	cfg := configs.GGBFlow

	sOpt := _grpc.ChainStreamInterceptor(interceptors.GetStreamValidateTokenInterceptor(cfg.APIKey, cfg.AuthorizedClient))
	cOpt := _grpc.ChainUnaryInterceptor(interceptors.GetUnaryValidateTokenInterceptor(cfg.APIKey, cfg.AuthorizedClient, "/ggbflow.GGBFlow/HealthCheck"))

	s = _grpc.NewServer(sOpt, cOpt)
	s.RegisterService(&gg_bflow.GGBFlow_ServiceDesc, f)

	return nil
}
