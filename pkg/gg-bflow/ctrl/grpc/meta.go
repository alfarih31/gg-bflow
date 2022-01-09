package grpc_ctrl

import (
	"context"
	"github.com/alfarih31/gg-bflow/api/grpc"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ctrl/grpc/errors"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/dto/meta"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/validator"
	grpc_message_adapters "github.com/alfarih31/gg-bflow/pkg/message-adapters"
	"github.com/alfarih31/gg-bflow/pkg/utils/datetime"
)

func (f *gRPCCtrl) SaveMeta(ctx context.Context, in *gg_bflow.MetaArg) (*gg_bflow.Ok, error) {
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

func (f *gRPCCtrl) LoadMeta(ctx context.Context, in *gg_bflow.LoadArg) (*gg_bflow.MetaRes, error) {
	res, err := f.metaSvc.Read(ctx, in.GetKey())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errors.ErrNotFound
	}

	return &gg_bflow.MetaRes{
		Key:       res.Key,
		Meta:      grpc_message_adapters.KeyValueToGRPCKeyValue(res.Metadata),
		CreatedAt: datetime.ToEpoch(res.CreatedAt),
		UpdatedAt: datetime.ToEpoch(res.UpdatedAt),
	}, nil
}
