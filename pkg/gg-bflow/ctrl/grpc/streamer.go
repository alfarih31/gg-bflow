package grpc_ctrl

import (
	"context"
	"github.com/alfarih31/gg-bflow/api/grpc"
	"github.com/alfarih31/gg-bflow/configs"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ctrl/grpc/errors"
	buffer_dto "github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/buffer"
	meta_dto "github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/svc/buffer"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/svc/meta"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/validator"
	grpc_message_adapters "github.com/alfarih31/gg-bflow/pkg/message-adapters"
	"io"
)

type streamerCtrl struct {
	gg_bflow.UnimplementedGGBFlowStreamerServer
	bufferSvc buffer.BufferSvc
	metaSvc   meta.MetaSvc
}

func (c streamerCtrl) SendDiscreteFlow(ctx context.Context, in *gg_bflow.SendFlowArg) (*gg_bflow.Ok, error) {
	arg := &buffer_dto.WriteArg{
		Key:  in.GetKey(),
		Data: in.GetByte(),
	}

	if err := validator.Validate(arg); err != nil {
		return nil, err
	}

	if configs.GGBFlow.BufferSizeLimit > 0 {
		if cap(arg.Data) > configs.GGBFlow.BufferSizeLimit {
			return nil, errors.ErrBufferLimitExceed
		}
	}

	err := c.bufferSvc.Write(ctx, *arg)
	if err != nil {
		return nil, err
	}

	return &gg_bflow.Ok{
		Message: "OK",
	}, nil
}

func (c streamerCtrl) SendFlow(s gg_bflow.GGBFlowStreamer_SendFlowServer) error {
	var res = new(gg_bflow.Ok)
	for {
		in, err := s.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		res, err = c.SendDiscreteFlow(s.Context(), in)
		if err != nil {
			return err
		}
	}

	return s.SendAndClose(res)
}

func (f *streamerCtrl) SaveMeta(ctx context.Context, in *gg_bflow.SaveMetaArg) (*gg_bflow.Ok, error) {
	arg := meta_dto.WriteArg{
		Key:  in.GetKey(),
		Meta: grpc_message_adapters.GRPCKeyValueToKeyValue(in.GetMeta()),
	}

	if err := validator.Validate(arg); err != nil {
		return nil, err
	}

	err := f.metaSvc.Write(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &gg_bflow.Ok{
		Message: "OK",
	}, nil
}
