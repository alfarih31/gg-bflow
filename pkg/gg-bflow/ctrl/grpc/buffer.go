package grpc_ctrl

import (
	"context"
	"github.com/alfarih31/gg-bflow/api/grpc"
	"github.com/alfarih31/gg-bflow/configs"
	buffer_svc "github.com/alfarih31/gg-bflow/pkg/gg-bflow/buffer/svc"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ctrl/grpc/errors"
	buffer_dto "github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/buffer"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/validator"
	"github.com/alfarih31/gg-bflow/pkg/message-adapters"
	keyvalue "github.com/alfarih31/nb-go-keyvalue"
	"io"
)

func (f *gRPCCtrl) SendDiscreteFlow(ctx context.Context, in *gg_bflow.SendArg) (*gg_bflow.Ok, error) {
	arg := &buffer_dto.WriteArg{
		Key:  in.GetKey(),
		Data: in.GetByteArr(),
	}

	if err := validator.Validate(arg); err != nil {
		return nil, err
	}

	if configs.GGBFlow.BufferSizeLimit > 0 {
		if cap(arg.Data) > configs.GGBFlow.BufferSizeLimit {
			return nil, errors.ErrBufferLimitExceed
		}
	}

	s, err := buffer_svc.Write(ctx, *arg)
	if err != nil {
		return nil, err
	}

	sKv, err := keyvalue.FromStruct(s)
	if err != nil {
		return nil, err
	}

	return &gg_bflow.Ok{
		Message: "OK",
		Meta:    message_adapters.KeyValueToGRPCKeyValue(sKv),
	}, nil
}

func (f *gRPCCtrl) LoadDiscreteFlow(ctx context.Context, in *gg_bflow.LoadArg) (*gg_bflow.FlowRes, error) {
	d, err := buffer_svc.Read(ctx, in.GetKey())

	if err != nil {
		return nil, err
	}

	if d == nil {
		return nil, errors.ErrNotFound
	}

	return &gg_bflow.FlowRes{
		ByteArr: d.Data,
	}, nil
}

func (f *gRPCCtrl) LoadFlow(s gg_bflow.GGBFlow_LoadFlowServer) error {
	for {
		in, err := s.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		res, err := f.LoadDiscreteFlow(s.Context(), in)
		if err != nil {
			return err
		}

		err = s.Send(res)
		if err != nil {
			return err
		}
	}
}

func (f *gRPCCtrl) SendFlow(s gg_bflow.GGBFlow_SendFlowServer) error {
	var res = new(gg_bflow.Ok)
	for {
		in, err := s.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		res, err = f.SendDiscreteFlow(s.Context(), in)
		if err != nil {
			return err
		}
	}

	return s.SendAndClose(res)
}
