package grpc_ctrl

import (
	"fmt"
	"github.com/alfarih31/gg-bflow/api/grpc"
	"github.com/alfarih31/gg-bflow/configs"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ctrl/grpc/interceptors"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/logger"
	_grpc "google.golang.org/grpc"
	"net"
	"time"
)

var startTime = time.Now()

var s *_grpc.Server

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

func Init() error {
	cfg := configs.GGBFlow

	sOpt := _grpc.ChainStreamInterceptor(interceptors.GetStreamValidateTokenInterceptor(cfg.APIKey, cfg.AuthorizedClient))
	cOpt := _grpc.ChainUnaryInterceptor(interceptors.GetUnaryValidateTokenInterceptor(cfg.APIKey, cfg.AuthorizedClient, "/ggbflow.GGBFlow/HealthCheck"))

	sCtrl := new(senderCtrl)

	vCtrl := new(loaderCtrl)

	s = _grpc.NewServer(sOpt, cOpt)
	s.RegisterService(&gg_bflow.GGBFlowSender_ServiceDesc, sCtrl)
	s.RegisterService(&gg_bflow.GGBFlowLoader_ServiceDesc, vCtrl)

	return nil
}
