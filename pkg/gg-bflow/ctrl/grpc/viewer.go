package grpc_ctrl

import (
	"context"
	"fmt"
	"github.com/alfarih31/gg-bflow/api/grpc"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ctrl/grpc/errors"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/svc/buffer"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/svc/meta"
	"github.com/alfarih31/gg-bflow/pkg/message-adapters"
	"github.com/alfarih31/gg-bflow/pkg/utils/datetime"
	"time"
)

type viewerCtrl struct {
	gg_bflow.UnimplementedGGBFlowViewerServer
	bufferSvc buffer.BufferSvc
	metaSvc   meta.MetaSvc
}

func (v *viewerCtrl) LoadDiscreteFlow(ctx context.Context, in *gg_bflow.LoadDiscreteFlowArg) (*gg_bflow.DiscreteFlowRes, error) {
	d, err := v.bufferSvc.Read(ctx, in.GetKey())

	if err != nil {
		return nil, err
	}

	if d == nil {
		return nil, errors.ErrNotFound
	}

	return &gg_bflow.DiscreteFlowRes{
		Byte: d.Data,
	}, nil
}

func (v *viewerCtrl) LoadFlow(a *gg_bflow.LoadFlowArg, s gg_bflow.GGBFlowViewer_LoadFlowServer) error {
	r := a.GetRate()
	info := ""
	ms := int64(time.Millisecond)
	if r == 0 {
		r = ms / 30
		info = "30 fps"
	} else {
		info = fmt.Sprintf("%d fps", r)
	}

	baseSleepMicroSec := ms / r

	for {
		initTime := time.Now().UnixMicro()

		res, err := v.LoadDiscreteFlow(s.Context(), &gg_bflow.LoadDiscreteFlowArg{
			Key: a.Key,
		})
		if err != nil {
			return err
		}

		err = s.Send(&gg_bflow.FlowRes{
			Byte: res.GetByte(),
			Info: info,
		})
		if err != nil {
			return err
		}

		time.Sleep(time.Duration(baseSleepMicroSec - (time.Now().UnixMicro() - initTime)))
	}
}

func (v *viewerCtrl) LoadMeta(ctx context.Context, in *gg_bflow.LoadMetaArg) (*gg_bflow.MetaRes, error) {
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
