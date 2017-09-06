package steps

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rai-project/pipeline"
)

type castToFloat32Slice struct {
	base
}

func NewCastToFloat32Slice(ctx context.Context) (pipeline.Step, error) {
	var res castToFloat32Slice
	return res.New(ctx)
}

func (p castToFloat32Slice) New(ctx context.Context) (pipeline.Step, error) {
	return p, nil
}

func (p castToFloat32Slice) Info() string {
	return "CastToFloat32Slice"
}

func (p castToFloat32Slice) do(ctx context.Context, in0 interface{}) interface{} {
	in, err := toSlice(in0)
	if err != nil {
		return errors.Errorf("expecting a slice input for CastToFloat32Slice, but got %v", in0)
	}
	res, err := toFloat32Slice(in)
	if err != nil {
		return err
	}
	return res
}

func (p castToFloat32Slice) Close() error {
	return nil
}

func init() {
	pipeline.Register(castToFloat32Slice{})
}
