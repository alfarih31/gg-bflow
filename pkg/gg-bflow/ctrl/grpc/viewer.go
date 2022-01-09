package grpc_ctrl

import (
	"context"
	"github.com/alfarih31/gg-bflow/api/grpc"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ctrl/grpc/errors"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/svc/buffer"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/svc/meta"
	"github.com/alfarih31/gg-bflow/pkg/message-adapters"
	"github.com/alfarih31/gg-bflow/pkg/utils/datetime"
	"io"
)

type viewerCtrl struct {
	gg_bflow.UnimplementedGGBFlowViewerServer
	bufferSvc buffer.BufferSvc
	metaSvc   meta.MetaSvc
}

func (v *viewerCtrl) LoadDiscreteFlow(ctx context.Context, in *gg_bflow.LoadArg) (*gg_bflow.FlowRes, error) {
	d, err := v.bufferSvc.Read(ctx, in.GetKey())

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

func (v *viewerCtrl) LoadFlow(s gg_bflow.GGBFlowViewer_LoadFlowServer) error {
	for {
		in, err := s.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		res, err := v.LoadDiscreteFlow(s.Context(), in)
		if err != nil {
			return err
		}

		err = s.Send(res)
		if err != nil {
			return err
		}
	}
}

func (v *viewerCtrl) LoadMeta(ctx context.Context, in *gg_bflow.LoadArg) (*gg_bflow.MetaRes, error) {
	res, err := v.metaSvc.Read(ctx, in.GetKey())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errors.ErrNotFound
	}

	return &gg_bflow.MetaRes{
		Key:       res.Key,
		Meta:      message_adapters.KeyValueToGRPCKeyValue(res.Metadata),
		CreatedAt: datetime.ToEpoch(res.CreatedAt),
		UpdatedAt: datetime.ToEpoch(res.UpdatedAt),
	}, nil
}
