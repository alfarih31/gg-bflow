package grpc_ctrl

import (
	"context"
	app "github.com/alfarih31/gg-bflow"
	gg_bflow "github.com/alfarih31/gg-bflow/api/grpc"
	grpc_message_adapters "github.com/alfarih31/gg-bflow/pkg/message-adapters"
	"github.com/alfarih31/nb-go-keyvalue"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type mainCtrl struct {
	gg_bflow.UnimplementedGGBFlowServer
}

func HealthCheck() keyvalue.KeyValue {
	u := time.Since(startTime).String()

	res := keyvalue.KeyValue{
		"uptime": u,
	}

	res.Assign(app.Meta, true)

	return res
}

func (m *mainCtrl) HealthCheck(ctx context.Context, in *emptypb.Empty) (*gg_bflow.HealthCheckRes, error) {
	s := HealthCheck()

	return &gg_bflow.HealthCheckRes{
		Data: grpc_message_adapters.KeyValueToGRPCKeyValue(s),
	}, nil
}
