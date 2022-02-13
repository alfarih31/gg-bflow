package grpc_ctrl

import (
	"context"
	"fmt"
	"github.com/alfarih31/gg-bflow/api/grpc"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/ctrl/grpc/errors"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/buffer"
	"github.com/alfarih31/gg-bflow/pkg/gg-bflow/repos/meta"
	"github.com/alfarih31/gg-bflow/pkg/utils"
	"time"
)

type loaderCtrl struct {
	gg_bflow.UnimplementedGGBFlowLoaderServer
}

func (v *loaderCtrl) LoadDiscreteFlow(ctx context.Context, in *gg_bflow.LoadDiscreteFlowArg) (*gg_bflow.DiscreteFlowRes, error) {
	d, err := buffer.Do.Read(ctx, in.GetKey())

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

func (v *loaderCtrl) LoadFlow(a *gg_bflow.LoadFlowArg, s gg_bflow.GGBFlowLoader_LoadFlowServer) error {
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

func (v *loaderCtrl) LoadMeta(ctx context.Context, in *gg_bflow.LoadMetaArg) (*gg_bflow.MetaRes, error) {
	res, err := meta.Do.Read(ctx, in.GetKey())
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, errors.ErrNotFound
	}

	return &gg_bflow.MetaRes{
		Key:       res.Key,
		Meta:      res.Metadata,
		CreatedAt: utils.NewDatetimeFromTime(res.CreatedAt).ToEpoch(),
		UpdatedAt: utils.NewDatetimeFromTime(res.UpdatedAt).ToEpoch(),
	}, nil
}
