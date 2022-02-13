package grpc_ctrl

import (
	"context"
	"github.com/alfarih31/gg-bflow/api/grpc"
	"github.com/alfarih31/gg-bflow/configs"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ctrl/grpc/errors"
	buffer_dto "github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/buffer"
	meta_dto "github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/buffer"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/meta"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/validator"
	"github.com/alfarih31/gg-bflow/pkg/utils"
	"io"
)

type senderCtrl struct {
	gg_bflow.UnimplementedGGBFlowSenderServer
}

func (c *senderCtrl) SendDiscreteFlow(ctx context.Context, in *gg_bflow.SendFlowArg) (*gg_bflow.Ok, error) {
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

	err := buffer.Do.Write(ctx, arg.Key, arg.Data)
	if err != nil {
		return nil, err
	}

	return &gg_bflow.Ok{
		Message: "OK",
	}, nil
}

func (c *senderCtrl) SendFlow(s gg_bflow.GGBFlowSender_SendFlowServer) error {
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

func (f *senderCtrl) SendMeta(ctx context.Context, in *gg_bflow.SendMetaArg) (*gg_bflow.Ok, error) {
	arg := meta_dto.WriteArg{
		Key:  in.GetKey(),
		Meta: in.GetMeta(),
	}

	if err := validator.Validate(arg); err != nil {
		return nil, err
	}

	// Get exist
	i, err := meta.Do.Read(ctx, arg.Key)
	if err != nil {
		return nil, err
	}

	now := utils.NewDatetimeNow()
	if i == nil {
		i = &meta_dto.Item{
			Key:      arg.Key,
			Metadata: arg.Meta,
		}

		err = meta.Do.Write(ctx, *i)
	} else {
		i.Metadata = arg.Meta
		i.UpdatedAt = now.GetTime()

		err = meta.Do.Update(ctx, arg.Key, *i)
	}

	if err != nil {
		return nil, err
	}

	return &gg_bflow.Ok{
		Message: "OK",
	}, nil
}
