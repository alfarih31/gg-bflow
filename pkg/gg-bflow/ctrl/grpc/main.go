package grpc_ctrl

import (
	"context"
	app "github.com/alfarih31/gg-bflow"
	gg_bflow "github.com/alfarih31/gg-bflow/api/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type mainCtrl struct {
	gg_bflow.UnimplementedGGBFlowServer
}

func (m *mainCtrl) HealthCheck(ctx context.Context, in *emptypb.Empty) (*gg_bflow.HealthCheckRes, error) {
	u := time.Since(startTime).String()

	return &gg_bflow.HealthCheckRes{
		Uptime: u,
		Meta:   app.Meta.String(),
	}, nil
}
