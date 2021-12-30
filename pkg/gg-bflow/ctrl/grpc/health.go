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

func HealthCheck(startTime time.Time, m keyvalue.KeyValue) keyvalue.KeyValue {
	u := time.Since(startTime).String()

	res := keyvalue.KeyValue{
		"uptime": u,
	}

	res.Assign(m, true)

	return res
}

func (f *gRPCCtrl) HealthCheck(ctx context.Context, in *emptypb.Empty) (*gg_bflow.HealthCheckRes, error) {
	kv, err := keyvalue.FromStruct(app.Meta)
	if err != nil {
		return nil, err
	}

	s := HealthCheck(f.st, kv)

	return &gg_bflow.HealthCheckRes{
		Data: grpc_message_adapters.KeyValueToGRPCKeyValue(s),
	}, nil
}
